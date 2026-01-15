package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/srthk29/grpc-example/proto/v2"
	"google.golang.org/grpc"
)

type propagationServer struct {
	pb.UnimplementedPropogationServiceServer
}

func (s *propagationServer) GetPropogation(ctx context.Context, req *pb.PropogationRequest) (*pb.PropogationReply, error) {
	log.Println(ctx)

	if req == nil {
		return nil, errors.New("nil request")
	}

	/*
		resp := &pb.PropogationReply{
			Tle: &pb.TLE{
				Name:  "ISS (ZARYA)",
				Line1: "1 25544U 98067A   26012.17690827  .00009276  00000+0  17471-3 0  9997",
				Line2: "2 25544  51.6333 351.7881 0007723   8.9804 351.1321 15.49250518547578",
			},
			Propogations: []*pb.Propogation{
				&pb.Propogation{
					Timestamp: &pb.Timestamp{
						Seconds: 1768223804,
					},
					LatLng: &pb.LatLng{
						Latitude:  -51.528,
						Longitude: -41.798,
					},
					AltitudeMeters: 437070.0,
				},
				&pb.Propogation{
					Timestamp: &pb.Timestamp{
						Seconds: 1768224104,
					},
					LatLng: &pb.LatLng{
						Latitude:  -50.957,
						Longitude: -35.962,
					},
					AltitudeMeters: 436492.0,
				},
				&pb.Propogation{
					Timestamp: &pb.Timestamp{
						Seconds: 1768224404,
					},
					LatLng: &pb.LatLng{
						Latitude:  -50.077,
						Longitude: -30.312,
					},
					AltitudeMeters: 435790.0,
				},
			},
		}
	*/

	resp := &pb.PropogationReply{
		Tle: &pb.Tle{
			Name:  "ISS (ZARYA)",
			Line1: "1 25544U 98067A   26012.17690827  .00009276  00000+0  17471-3 0  9997",
			Line2: "2 25544  51.6333 351.7881 0007723   8.9804 351.1321 15.49250518547578",
			Age:   0.42,
			Propogation: &pb.Propogation{
				Timestamp: 1768223804,
				Latitude:  -51.528,
				Longitude: -41.798,
				Altitude:  437070.0,
			},
		},
		Propogations: []*pb.Propogation{
			&pb.Propogation{
				Timestamp: 1768223804,
				Latitude:  -51.528,
				Longitude: -41.798,
				Altitude:  437070.0,
			},
			&pb.Propogation{
				Timestamp: 1768224104,
				Latitude:  -50.957,
				Longitude: -35.962,
				Altitude:  436492.0,
			},
			&pb.Propogation{
				Timestamp: 1768224404,
				Latitude:  -50.077,
				Longitude: -30.312,
				Altitude:  435790.0,
			},
		},
	}

	return resp, nil
}

// https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go
func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterPropogationServiceServer(grpcServer, &propagationServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
