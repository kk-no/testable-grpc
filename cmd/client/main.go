package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/kk-no/proto-terminal/sample/v1"
	"github.com/kk-no/testable-grpc/grpc/client"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to dial for grpc", err)
	}
	defer conn.Close()

	sc := pb.NewSampleServiceClient(conn)

	cli := client.NewSampleClient(sc)
	cli.Ping(ctx)
}
