package main

import (
	"hackaton/api"
	"hackaton/storage"
)

func main() {

	storage.ConnectStorage()

	server := api.APIServer{Addr: ":8000"}
	server.Run()

}
