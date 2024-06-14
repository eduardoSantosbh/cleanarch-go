package order

import (
	"github.com/codesantos/cleanarch/pkg"
	"time"
)

type Order struct {
	ID         pkg.ID
	Price      float64
	Tax        float64
	FinalPrice float64
	CreatedAt  time.Time
}

func NewOrder(price float64, tax float64, dateOrder time.Time) (*Order, error) {
	order := &Order{
		ID:        pkg.NewID(),
		Price:     price,
		Tax:       tax,
		CreatedAt: dateOrder,
	}
	err := order.isValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) create(price float64, tax float64) {
	o.Price = price
	o.Tax = tax
}

func (o *Order) isValid() error {
	validators := []Validator{
		&IDValidator{},
		&PriceValidator{},
		&TaxValidator{},
	}
	for _, validator := range validators {
		if err := validator.Validate(o); err != nil {
			return err
			break
		}
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.isValid()
	if err != nil {
		return err
	}
	return nil
}
