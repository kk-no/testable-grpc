package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/kk-no/proto-terminal/sample/v1"
	"github.com/kk-no/testable-grpc/grpc/server"
)

func main() {
	port, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("Failed to listen server:", err)
	}
	s := grpc.NewServer()
	pb.RegisterSampleServiceServer(s, server.NewSampleServer())

	if err := s.Serve(port); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
