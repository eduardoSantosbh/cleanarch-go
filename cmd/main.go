package main

import (
	"fmt"
	"github.com/codesantos/cleanarch/internal/infra/database"
	"github.com/codesantos/cleanarch/internal/infra/repository"
	"github.com/codesantos/cleanarch/internal/infra/web"
	"log"
)

func main() {
	adapter, err := database.NewMySQLAdapter()
	if err != nil {
		log.Fatalf("Failed to create MySQL adapter: %v", err)
	}
	repo := repository.NewOrderRepositoryImpl(*adapter)
	webOrderHAndler := web.NewOrderHandler(repo)

	weberver := web.NewWebServer(":8000")
	weberver.AddHandler("/v1/order", webOrderHAndler.Create)
	fmt.Println("Server running on port 8000")
	weberver.Start()
}
