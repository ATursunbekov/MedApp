package grpc

import (
	pb "github.com/ATursunbekov/MedApp/proto"
	"google.golang.org/grpc"
)

func NewAnamnesisClient(addr string) (pb.SessionServiceClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return pb.NewSessionServiceClient(conn), nil
}
