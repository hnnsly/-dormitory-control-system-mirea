package main

import (
	"hackaton/api"
	"hackaton/storage"
)

func main() {

	storage.ConnectStorage()
	//_, err := storage.Store.Db.Exec("INSERT INTO residences VALUES ($1, $2, $3, $4, $5,$6)", "112", "2", "3", "1", "1", "0")
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//storage.Store.Db.Exec("INSERT INTO residences VALUES ($1, $2, $3, $4, $5,$6)", "13", "2", "3", "1", "2", "0")
	//storage.Store.Db.Exec("INSERT INTO residences VALUES ($1, $2, $3, $4, $5,$6)", "14", "2", "3", "1", "3", "0")
	//storage.Store.Db.Exec("INSERT INTO residences VALUES ($1, $2, $3, $4, $5,$6)", "15", "2", "3", "1", "4", "0")
	//storage.Store.Db.Exec("INSERT INTO residences VALUES ($1, $2, $3, $4, $5,$6)", "16", "2", "3", "1", "5", "0")
	//storage.Store.Db.Exec("INSERT INTO residences VALUES ($1, $2, $3, $4, $5,$6)", "1", "2", "3", "1", "6", "0")
	//storage.Store.Db.Exec("INSERT INTO residences VALUES ($1, $2, $3, $4, $5,$6)", "2", "2", "3", "1", "7", "0")
	//storage.Store.Db.Exec("INSERT INTO residences VALUES ($1, $2, $3, $4, $5,$6)", "3", "2", "3", "1", "8", "0")
	//storage.Store.Db.Exec("INSERT INTO residences VALUES ($1, $2, $3, $4, $5,$6)", "4", "2", "3", "1", "9", "0")
	//storage.Store.Db.Exec("INSERT INTO residences VALUES ($1, $2, $3, $4, $5,$6)", "5", "2", "3", "1", "10", "0")

	server := api.APIServer{Addr: ":8001"}
	server.Run()
}
