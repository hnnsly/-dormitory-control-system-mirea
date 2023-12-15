package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"math/rand"
)

var Store PStorage

func ConnectStorage() {

	Store.createStorage()

	err := Store.InitResidences()
	if err != nil {
		panic(fmt.Sprintf("could not fill 'Общага' table: %v", err))
	}

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
    		residence_address VARCHAR(255),
    		residence_id INT
		);
	`)

	_, err = store.Db.Exec(`CREATE TABLE IF NOT EXISTS residences (
			id SERIAL PRIMARY KEY,
			address VARCHAR(255),
			floor INTEGER,
			room INTEGER,
			place INTEGER,
			is_occupied INTEGER
		)
`)
	if err != nil {
		panic(fmt.Sprintf("could not create 'students' table: %v", err))
	}

}

type PStorage struct { // PStorage - PostgreSQL Store
	Db *sql.DB
}

func (store *PStorage) InitResidences() error {
	var rowCount int
	err := store.Db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", residences)).Scan(&rowCount)
	if rowCount != 0 {
		return nil
	}
	addresses := []string{"Ул. Пушкина, д. 21", "Ул. Вернадского, д.86к4", "Ул. Асанова, д.14к8"}
	c := 0
	for _, address := range addresses {
		for floor := 1; floor <= rand.Intn(3)+4; floor++ {
			for room := 1; room <= rand.Intn(120)+20; room++ {
				lim := 0
				if room%2 == 0 {
					lim = 3
				} else {
					lim = 2
				}
				for place := 1; place <= lim; place++ {
					c++
					_, err := store.Db.Exec("INSERT INTO residences VALUES ($1, $2, $3, $4, $5,$6)", c, address, floor, room, place, 0)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}
