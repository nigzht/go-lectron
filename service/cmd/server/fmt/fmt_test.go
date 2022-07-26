package fmt_test

import (
	"context"
	"encoding/json"
	"testing"
	fm "service/cmd/server/fmt"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestPrint(t *testing.T) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:6969", opts...)
	if err != nil {
		t.Error(err)
	}

	defer conn.Close()

	c := fm.NewFmtClient(conn)

	c.Print(context.Background(), &anypb.Any{
		TypeUrl: "type.googleapis.com/google.protobuf.StringValue",
		Value:   []byte("Hello World"),
	})

	m, _ := json.Marshal(map[string]interface{}{
		"ping": "pong",
	})

	c.Print(context.Background(), &anypb.Any{
		TypeUrl: "type.googleapis.com/google.protobuf.StringValue",
		Value:   m,
	})
}
