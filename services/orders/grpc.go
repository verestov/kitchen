package main

import (
	"fmt"
	"log"
	"net"

	handler "github.com/verestov/kitchen/services/orders/handler/orders"
	"github.com/verestov/kitchen/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()

	// we need to register our gRPC servicies
	orderService := service.NewOrderService()
	handler.NewGrpcOrderService(grpcServer, orderService)

	fmt.Println("Strting server on ", s.addr)

	return grpcServer.Serve(lis)
}
