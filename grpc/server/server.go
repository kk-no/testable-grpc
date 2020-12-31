package server

import (
	"context"

	pb "github.com/kk-no/proto-terminal/sample/v1"
)

type SampleServer struct{}

func NewSampleServer() pb.SampleServiceServer {
	return &SampleServer{}
}

func (s *SampleServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	println("received ping request")
	return &pb.PingResponse{}, nil
}
