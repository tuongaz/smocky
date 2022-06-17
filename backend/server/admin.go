package server

import (
	"context"
	"fmt"
)

type AdminServer struct{}

func NewAdminServer() *AdminServer {
	return &AdminServer{}
}

func (*AdminServer) Start(ctx context.Context, port int32) (string, func(), error) {
	return "http://localhost:2601", func() {
		fmt.Println("shutting down admin server")
	}, nil
}
