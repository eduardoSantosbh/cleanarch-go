package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const dns = "root:root@tcp(127.0.0.1:3306)/orders_db"

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
