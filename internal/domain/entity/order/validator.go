package order

import (
	"errors"
	"github.com/codesantos/cleanarch/pkg"
)

var (
	ErrInvalidID    = errors.New("invalid ID")
	ErrInvalidPrice = errors.New("invalid price")
	ErrInvalidTax   = errors.New("invalid tax")
)

type Validator interface {
	Validate(*Order) error
}

type IDValidator struct{}

func (v *IDValidator) Validate(o *Order) error {
	if _, err := pkg.ParseID(o.ID.String()); err != nil {
		return ErrInvalidID
	}
	return nil
}

type PriceValidator struct{}

func (v *PriceValidator) Validate(o *Order) error {
	if o.Price <= 0 {
		return ErrInvalidPrice
	}
	return nil
}

type TaxValidator struct{}

func (v *TaxValidator) Validate(o *Order) error {
	if o.Tax <= 0 {
		return ErrInvalidTax
	}
	return nil
}
