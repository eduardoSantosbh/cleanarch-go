package web

import (
	"encoding/json"
	"github.com/codesantos/cleanarch/internal/application/usecase"
	"github.com/codesantos/cleanarch/internal/infra/repository"
	"net/http"
)

type OrderHandler struct {
	OrderRepository repository.OrderRepository
}

func NewOrderHandler(repository repository.OrderRepository) *OrderHandler {
	return &OrderHandler{OrderRepository: repository}
}

func (repository *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createOrder := usecase.CreateOrderUseCase{Repository: repository.OrderRepository}
	output, err := createOrder.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
