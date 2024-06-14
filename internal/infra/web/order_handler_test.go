package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/codesantos/cleanarch/internal/application/usecase"
	"github.com/codesantos/cleanarch/internal/domain/entity/order"
	"github.com/codesantos/cleanarch/internal/infra/database"
	"github.com/codesantos/cleanarch/internal/infra/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Save(order *order.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func GivenValidInput_WhenCreate_ThenStatusCreated(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	mockRepo.On("Save", mock.Anything).Return(nil)

	adapter, err := database.NewMySQLAdapter()
	if err != nil {
		log.Fatalf("Failed to create MySQL adapter: %v", err)
	}

	repo := repository.NewOrderRepositoryImpl(*adapter)

	handler := NewOrderHandler(repo)
	//handler := NewOrderHandler(mockRepo)

	orderInput := usecase.Input{Price: 100.0, Tax: 10.0}
	orderInputJson, _ := json.Marshal(orderInput)

	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(orderInputJson))
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockRepo.AssertExpectations(t)
}

func GivenRepoError_WhenCreate_ThenStatusInternalServerError(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	mockRepo.On("Save", mock.Anything).Return(errors.New("error"))

	handler := NewOrderHandler(mockRepo)

	orderInput := usecase.Input{Price: 100.0, Tax: 10.0}
	orderInputJson, _ := json.Marshal(orderInput)

	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(orderInputJson))
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockRepo.AssertExpectations(t)
}

func GivenInvalidInput_WhenCreate_ThenStatusBadRequest(t *testing.T) {
	mockRepo := new(MockOrderRepository)

	handler := NewOrderHandler(mockRepo)

	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer([]byte("invalid")))
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	mockRepo.AssertExpectations(t)
}
