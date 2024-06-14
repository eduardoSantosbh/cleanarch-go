package usecase

import (
	"github.com/codesantos/cleanarch/internal/domain/entity/order"
	"github.com/codesantos/cleanarch/internal/infra/repository"
	"time"
)

type Input struct {
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type Output struct {
	ID string `json:"id"`
}

type CreateOrderUseCase struct {
	Repository repository.OrderRepository
}

func (c *CreateOrderUseCase) Execute(input Input) (Output, error) {
	order, err := order.NewOrder(input.Price, input.Tax, time.Time{})
	if err != nil {
		return Output{}, err
	}
	order.CalculateFinalPrice()
	if err := c.Repository.Save(order); err != nil {
		return Output{}, err
	}
	output := Output{
		ID: order.ID.String(),
	}
	return output, nil
}
