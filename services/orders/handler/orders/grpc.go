//TODO: https://www.youtube.com/watch?v=ea_4Ug5WWYE&t=32s&ab_channel=Tiago
//FIXME: ±20:31 time

package handler

import (
	"context"

	"github.com/verestov/kitchen/services/common/genproto/orders"
	"github.com/verestov/kitchen/services/orders/types"
)

type OrderGrpcHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrderService() *OrderGrpcHandler {
	gRPCHandler := &OrderGrpcHandler{}

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
