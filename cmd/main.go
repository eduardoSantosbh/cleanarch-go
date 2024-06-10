package main

import (
	"fmt"
	"github.com/codesantos/cleanarch/internal/application/usecase"
	"github.com/codesantos/cleanarch/internal/infra/database"
	"github.com/codesantos/cleanarch/internal/infra/repository"
	"log"
)

func main() {

	adapter, err := database.NewMySQLAdapter()
	if err != nil {
		log.Fatalf("Failed to create MySQL adapter: %v", err)
	}
	// 1. Criar uma instância do repositório
	repo := repository.NewOrderRepositoryImpl(*adapter)

	// 2. Criar uma instância do caso de uso
	createOrderUseCase := &usecase.CreateOrderUseCase{
		Repository: repo,
	}

	// 3. Criar uma entrada (input) para o caso de uso
	input := usecase.Input{
		Price: 100,
		Tax:   10.5,
	}

	// 4. Executar o caso de uso
	output, err := createOrderUseCase.Execute(input)
	if err != nil {
		log.Fatalf("Failed to create order: %v", err)
	}

	// Exibir o resultado
	fmt.Printf("Order created successfully!\nID: %s\nPrice: %.2f\nTax: %.2f\nFinal Price: %.2f\n",
		output.ID, output.Price, output.Tax, output.FinalPrice)

}
