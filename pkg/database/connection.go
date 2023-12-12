package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "hackaton/pkg/models"
)

var DB *sql.DB

func Connect() {
	connectionString := "host=localhost port=5432 user=postgres password=testtest dbname=dormitory sslmode=disable"
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
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			fio VARCHAR(255) NOT NULL,
			date_of_birth VARCHAR(10) NOT NULL
		);
	`)
	if err != nil {
		panic(fmt.Sprintf("could not create 'students' table: %v", err))
	}
}
