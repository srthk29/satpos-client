package main

import (
	"context"
	"log"
	"time"

	pb "github.com/srthk29/grpc-example/proto/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getPropogation(ctx context.Context, client pb.PropogationServiceClient, noradCatalog int32) {
	log.Printf("Getting propation for NORAD (%d)", noradCatalog)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetPropogation(ctx, &pb.PropogationRequest{NoradCategory: noradCatalog})
	if err != nil {
		log.Fatalf("client.GetPropogation failed: %v", err)
	}

	log.Println(resp)
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

	client := pb.NewPropogationServiceClient(conn)

	getPropogation(context.Background(), client, 25544)
}
