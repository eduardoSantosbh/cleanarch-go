package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const dns = "root:root@tcp(localhost:3306)/orders_db?charset=utf8mb4&parseTime=True&loc=Local"

type AdapterMysql struct {
	Connection *sql.DB
}

func NewMySQLAdapter() (*AdapterMysql, error) {
	db, err := sql.Open("mysql", dns)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return &AdapterMysql{Connection: db}, nil
}
