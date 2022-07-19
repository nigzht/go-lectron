package server

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthenticationServerImpl struct{}

func NewAuthenticationServerImpl() *AuthenticationServerImpl {
	return &AuthenticationServerImpl{}
}

func (a *AuthenticationServerImpl) Login(context.Context, *emptypb.Empty) (*LoginDetails, error) {
	panic(fmt.Errorf("unimplemented"))
}

func (a *AuthenticationServerImpl) mustEmbedUnimplementedAuthenticationServer() {}
