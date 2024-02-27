package usecase

import (
	"github.com/carloseduribeiro/order-service-clean-arch/internal/domain/entity"
)

type ListOrdersUseCase struct {
	orderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{orderRepository: orderRepository}
}

type ListOrdersOutputDto struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

func (c *ListOrdersUseCase) Execute() ([]ListOrdersOutputDto, error) {
	orders, err := c.orderRepository.FindAll()
	if err != nil {
		return nil, err
	}
	result := make([]ListOrdersOutputDto, 0, len(orders))
	for _, order := range orders {
		result = append(result, ListOrdersOutputDto{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return result, nil
}
