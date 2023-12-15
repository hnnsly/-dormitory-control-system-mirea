package main

import (
	"hackaton/api"
	"hackaton/storage"
)

func main() {

	storage.ConnectStorage()
	stud, _ := storage.Store.ShowStudentsByCriteria("full_name", "Игорь Петрович Тортуга", 0)
	storage.Store.Add(&stud[0][0])
	storage.Store.Add(&stud[0][0])
	storage.Store.Add(&stud[0][0])
	storage.Store.Add(&stud[0][0])
	storage.Store.Add(&stud[0][0])

	server := api.APIServer{Addr: ":8000"}
	server.Run()
}
