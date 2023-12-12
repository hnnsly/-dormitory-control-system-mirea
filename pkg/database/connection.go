package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "hackaton/pkg/models"
)

var DB *sql.DB

func Connect() {
	connectionString := "host=localhost port=5432 user=postgres password=7595 dbname=dormitory sslmode=disable"
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic("could not connect to the database")
	}

	DB = db

	err = DB.Ping()
	if err != nil {
		panic(fmt.Sprintf("could not ping the database: %v", err))
	}

	createTables()
}

func createTables() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL
		);
	`)
	if err != nil {
		panic(fmt.Sprintf("could not create 'users' table: %v", err))
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS students (
			card_number INT PRIMARY KEY,
    		full_name VARCHAR(255),
    		birth_date DATE,
    		photo_url VARCHAR(255),
    		housing_order_number INT,
    		enrollment_order_number INT,
    		enrollment_date DATE,
    		birth_place VARCHAR(255),
    		residence_address VARCHAR(255)
		);
	`)
	if err != nil {
		panic(fmt.Sprintf("could not create 'students' table: %v", err))
	}
}
