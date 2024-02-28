package web

import (
	"encoding/json"
	"net/http"

	"github.com/carloseduribeiro/order-service-clean-arch/internal/application/usecase"
	"github.com/carloseduribeiro/order-service-clean-arch/internal/domain/entity"
)

type ListOrdersHandler struct {
	orderRepository entity.OrderRepositoryInterface
}

func NewListOrdersHandler(orderRepository entity.OrderRepositoryInterface) *ListOrdersHandler {
	return &ListOrdersHandler{orderRepository: orderRepository}
}

func (l *ListOrdersHandler) List(w http.ResponseWriter, r *http.Request) {
	listOrders := usecase.NewListOrdersUseCase(l.orderRepository)
	output, err := listOrders.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
