package server

import (
	"fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

// Start starts the GCP server
func Start(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	fmt.Printf("server listening on port %d \n", port)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	return grpcServer.Serve(lis)
}
