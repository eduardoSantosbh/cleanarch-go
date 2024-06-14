package usecase

import (
	"errors"
	"github.com/codesantos/cleanarch/internal/domain/entity/order"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Save(order *order.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func TestGivenOrder_WhenSave_ThenSuccess(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	mockRepo.On("Save", mock.Anything).Return(nil)

	createOrder := CreateOrderUseCase{Repository: mockRepo}
	input := Input{Price: 10.0, Tax: 2.0}

	output, err := createOrder.Execute(input)

	assert.NoError(t, err)
	assert.NotNil(t, output.ID)

}

func TestExecuteFailsWhenOrderSavingFails(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	mockRepo.On("Save", mock.Anything).Return(errors.New("save error"))

	createOrder := CreateOrderUseCase{Repository: mockRepo}
	input := Input{Price: 100.0, Tax: 10.0}
	_, err := createOrder.Execute(input)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
