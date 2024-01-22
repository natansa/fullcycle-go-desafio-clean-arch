package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type ListOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrderUseCase) Execute() ([]ListOrderOutputDTO, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return []ListOrderOutputDTO{}, err
	}

	var dtos []ListOrderOutputDTO
	for _, order := range orders {
		dtos = append(dtos, ListOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		})
	}

	return dtos, nil
}
