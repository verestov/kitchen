package handler

import (
	"net/http"

	"github.com/verestov/kitchen/services/common/genproto/orders"
	"github.com/verestov/kitchen/services/common/util"
	"github.com/verestov/kitchen/services/orders/types"
)

type OrderHttpHandler struct {
	orderService types.OrderService
}

func NewHttpOrderHandler(orderService types.OrderService) *OrderHttpHandler {
	handler := &OrderHttpHandler{
		orderService: orderService,
	}

	return handler
}

func (h *OrderHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrderHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = h.orderService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	util.WriteJSON(w, http.StatusOK, res)
}
