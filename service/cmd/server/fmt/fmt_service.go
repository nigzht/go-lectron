package fmt

import (
	context "context"
	"fmt"

	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type FmtServerImpl struct {
	UnimplementedFmtServer
}

func NewFmtServer() *FmtServerImpl {
	return &FmtServerImpl{}
}

func (f *FmtServerImpl) Print(_ context.Context, d *anypb.Any) (*emptypb.Empty, error) {
	fmt.Print(d)
	return &emptypb.Empty{}, nil
}
