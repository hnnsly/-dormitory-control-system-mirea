package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Store PStorage

func ConnectStorage() {

	Store.createStorage()
	Store.initTables()

}

func (store *PStorage) createStorage() {
	connectionString := "host=localhost port=5432 user=postgres password=7595 dbname=dormitory sslmode=disable"
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic("could not connect to the database")
	}

	store.Db = db

	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("could not ping the database: %v", err))
	}

	store.initTables()
}

func (store *PStorage) initTables() {
	_, err := store.Db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL
		);
	`)
	if err != nil {
		panic(fmt.Sprintf("could not create 'users' table: %v", err))
	}
	//
	_, err = store.Db.Exec(`
		CREATE TABLE IF NOT EXISTS students (
		    id SERIAL PRIMARY KEY,
			card_number VARCHAR(255),
    		full_name VARCHAR(255),
    		birth_date VARCHAR(255),
    		photo_url TEXT,
    		housing_order_number VARCHAR(255),
    		enrollment_order_number VARCHAR(255),
    		enrollment_date VARCHAR(255),
    		birth_place VARCHAR(255),
    		residence_address VARCHAR(255)
		);
	`)
	if err != nil {
		panic(fmt.Sprintf("could not create 'students' table: %v", err))
	}
}

type PStorage struct { // PStorage - PostgreSQL Store
	Db *sql.DB
}
