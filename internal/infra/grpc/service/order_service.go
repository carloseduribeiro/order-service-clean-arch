package service

import (
	"context"

	"github.com/carloseduribeiro/order-service-clean-arch/internal/application/usecase"
	"github.com/carloseduribeiro/order-service-clean-arch/internal/infra/grpc/pb"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrderService(
	createOrderUseCase usecase.CreateOrderUseCase,
	listOrdersUseCase usecase.ListOrdersUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.Order, error) {
	dto := usecase.OrderInputDto{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.Order{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(context.Context, *pb.Blank) (*pb.OrderListResponse, error) {
	orders, err := s.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	response := make([]*pb.Order, 0, len(orders))
	for _, order := range orders {
		response = append(response, &pb.Order{Id: order.ID, Price: float32(order.Price), Tax: float32(order.Tax)})
	}
	return &pb.OrderListResponse{Orders: response}, nil
}
