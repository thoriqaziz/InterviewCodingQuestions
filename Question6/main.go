// Database pooling

package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	conn := "host=localhost port=5432 user=postgres password=Shahia01 dbname=restapi sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	DB.SetMaxOpenConns(35)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(30 * time.Minute)
	DB.SetConnMaxIdleTime(10 * time.Minute)

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected with pooling")
}

func main() {
	Connect()
}
