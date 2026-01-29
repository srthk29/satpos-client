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

func plot(ctx context.Context, propagations []*pb.Propagation) {
	sat := model.SatellitePlot{
		PlotTypes:   []model.PlotType{model.PlotTypePlateCarree},
		Size:        model.SizePrint,
		Format:      model.ImageFormatSVG,
		Colorscheme: model.ColorschemeDefault,
	}

	locations := make([]*model.Location, 0, len(propagations))
	for _, prop := range propagations {
		locations = append(locations, &model.Location{
			Latitude:  prop.Geodetic.LatitudeDeg,
			Longitude: prop.Geodetic.LongitudeDeg,
			Altitude:  float32(prop.Geodetic.AltitudeKm),
		})
	}
	sat.Locations = locations

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

func getPropogation(ctx context.Context, client pb.PropagationServiceClient, noradCatalog int32) ([]*pb.Propagation, error) {
	log.Printf("Getting propation for NORAD (%d)", noradCatalog)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetPropagation(ctx, &pb.GetPropagationRequest{SatelliteNumber: uint32(noradCatalog)})
	if err != nil {
		//log.Fatalf("client.GetPropogation failed: %v", err)
		return nil, err
	}

	log.Println(resp)

	return resp.Propagations, nil
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

	propogations, err := getPropogation(ctx, client, 25544)
	if err != nil {
		log.Fatalf("client.GetPropogation failed: %v", err)
	}

	plot(ctx, propogations)
}
