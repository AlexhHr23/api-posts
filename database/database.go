package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(dsn string) error {
	var err error
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		return fmt.Errorf("Error al abrir la base de datos: %w", err)
	}

	err = DB.Ping()

	if err != nil {
		return fmt.Errorf("Error al abrir la base de datos: %w", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	return nil
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
