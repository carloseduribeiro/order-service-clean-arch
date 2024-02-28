package web

import (
	"encoding/json"
	"net/http"

	"github.com/carloseduribeiro/order-service-clean-arch/internal/application/usecase"
	"github.com/carloseduribeiro/order-service-clean-arch/internal/domain/entity"
	"github.com/carloseduribeiro/order-service-clean-arch/pkg/events"
)

type CreateOrderHttpHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
) *CreateOrderHttpHandler {
	return &CreateOrderHttpHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}

func (h *CreateOrderHttpHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
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
