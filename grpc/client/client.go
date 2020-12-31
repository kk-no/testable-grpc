package client

import (
	"context"
	"log"

	pb "github.com/kk-no/proto-terminal/sample/v1"
)

type SampleClientManager interface {
	Ping(ctx context.Context)
}

type SampleClient struct {
	grpc pb.SampleServiceClient
}

func (s *SampleClient) Ping(ctx context.Context) {
	res, err := s.grpc.Ping(ctx, &pb.PingRequest{})
	if err != nil {
		log.Fatal("Failed to sample ping", err)
	}

	println(res.String())
}

func NewSampleClient(cli pb.SampleServiceClient) SampleClientManager {
	return &SampleClient{grpc: cli}
}
