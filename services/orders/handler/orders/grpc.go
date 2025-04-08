package handler

import (
	"context"

	"github.com/verestov/kitchen/services/common/genproto/orders"
	"github.com/verestov/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

type OrderGrpcHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrderService(grpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrderGrpcHandler{
		ordersService: ordersService,
	}

	orders.RegisterOrderServiceServer(grpc, gRPCHandler)

}

func (h *OrderGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  3,
		Quantity:   5,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}
