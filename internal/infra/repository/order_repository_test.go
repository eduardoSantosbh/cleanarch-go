package repository

import (
	"database/sql"
	"testing"

	"github.com/codesantos/cleanarch/internal/domain/entity/order"
	"github.com/codesantos/cleanarch/internal/infra/database"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open SQLite database: %v", err)
	}

	createTableQuery := `
	CREATE TABLE orders (
		id TEXT PRIMARY KEY,
		price REAL,
		tax REAL,
		final_price REAL
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		t.Fatalf("Failed to create orders table: %v", err)
	}

	return db
}

func TestOrderRepositoryImpl_GivenValidOrder_WhenSaveIsCalled_ThenOrderIsInsertedSuccessfully(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	adapter := database.AdapterMysql{Connection: db}
	repo := NewOrderRepositoryImpl(adapter)

	expectedOrder := &order.Order{
		Price:      100.0,
		Tax:        10.0,
		FinalPrice: 110.0,
	}

	err := repo.Save(expectedOrder)
	assert.NoError(t, err)

	row := db.QueryRow("SELECT price, tax, final_price FROM orders WHERE id = ?", expectedOrder.ID)
	var price, tax, finalPrice float64
	err = row.Scan(&price, &tax, &finalPrice)
	assert.NoError(t, err)
	assert.Equal(t, expectedOrder.Price, price)
	assert.Equal(t, expectedOrder.Tax, tax)
	assert.Equal(t, expectedOrder.FinalPrice, finalPrice)
}
