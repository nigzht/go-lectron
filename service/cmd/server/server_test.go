package server_test

import (
	"testing"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestConn(t *testing.T) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:6969", opts...)
	if err != nil {
		t.Errorf("%v", err)
	}
	defer conn.Close()
}
