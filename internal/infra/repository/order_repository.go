package repository

import (
	"github.com/codesantos/cleanarch/internal/domain/entity/order"
	"github.com/codesantos/cleanarch/internal/infra/database"
)

type OrderRepository interface {
	Save(order *order.Order) error
}

type OrderRepositoryImpl struct {
	Adapter database.AdapterMysql
}

func NewOrderRepositoryImpl(adapter database.AdapterMysql) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{Adapter: adapter}
}

func (repo *OrderRepositoryImpl) Save(order *order.Order) error {
	stmt, err := repo.Adapter.Connection.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}
