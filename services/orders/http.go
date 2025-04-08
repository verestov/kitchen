package main

import (
	"log"
	"net/http"

	handler "github.com/verestov/kitchen/services/orders/handler/orders"
	"github.com/verestov/kitchen/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (h *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	handler.NewHttpOrderHandler(orderService)

	log.Println("Starting server on:", h.addr)

	return http.ListenAndServe(h.addr, router)
}
