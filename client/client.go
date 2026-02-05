package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/srthk29/grpc-example/model"
	pb "github.com/srthk29/grpc-example/proto/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const NoradId int64 = 25544

func plot(ctx context.Context, resppb *pb.GetPropagationResponse) {
	sat := model.SatellitePlot{
		NoradId:   NoradId,
		PlotTypes: []model.PlotType{model.PlotTypePlateCarree},
		Size:      model.SizeMedium,
		Format:    model.ImageFormatSVG,
		Colorscheme: model.Colorscheme{
			Accent: model.AccentPurple,
			Theme:  model.ThemeMuted,
		},
		ShowIcon:      true,
		AddNightShade: true,
		Features:      []model.Feature{},
	}

	locations := make([]*model.GeodeticPosition, 0, len(resppb.Propagations))
	for _, prop := range resppb.Propagations {
		locations = append(locations, &model.GeodeticPosition{
			Latitude:   prop.Geodetic.LatitudeDeg,
			Longtitude: prop.Geodetic.LongitudeDeg,
			Altitude:   prop.Geodetic.AltitudeKm,
		})
	}
	sat.Positions = locations

	sat.NowPosition = &model.GeodeticPosition{
		Latitude:   resppb.AtNowUtc.Geodetic.LatitudeDeg,
		Longtitude: resppb.AtNowUtc.Geodetic.LongitudeDeg,
		Altitude:   resppb.AtNowUtc.Geodetic.AltitudeKm,
	}

	payload, err := json.Marshal(sat)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"http://127.0.0.1:8000/custom/plot",
		bytes.NewReader(payload))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		panic(fmt.Errorf("request failed: %s\n%s", resp.Status, body))
	}

	fileName := "plot_" + strconv.Itoa(int(time.Now().UTC().Unix())) + ".svg"
	out, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		panic(err)
	}
}

func getPropogation(ctx context.Context, client pb.PropagationServiceClient, noradCatalog int64) (*pb.GetPropagationResponse, error) {
	log.Printf("Getting propation for NORAD (%d)", noradCatalog)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetPropagation(ctx, &pb.GetPropagationRequest{SatelliteNumber: uint32(noradCatalog)})
	if err != nil {
		//log.Fatalf("client.GetPropogation failed: %v", err)
		return nil, err
	}

	log.Println(resp)

	return resp, nil
}

// https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewPropagationServiceClient(conn)

	ctx := context.Background()

	propogations, err := getPropogation(ctx, client, NoradId)
	if err != nil {
		log.Fatalf("client.GetPropogation failed: %v", err)
	}

	plot(ctx, propogations)
}
