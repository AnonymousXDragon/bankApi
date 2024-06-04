package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func MakeConnection() *sql.DB {
	driver := "postgres"
	dbUrl := "postgresql://postgres:password@localhost:5432/bankdb?sslmode=disable"

	conn, err := sql.Open(driver, dbUrl)

	if err != nil {
		fmt.Println("error", err)
		log.Fatal(err.Error())
	}

	return conn
}
