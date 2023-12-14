package main

import (
	"hackaton/api"
	"hackaton/storage"
)

func main() {
	//database.ConnectStorage()
	//databaseModels.InitStudentsDB()
	//cache, err := templates.NewTemplateCache("web/html/")
	//if err != nil {
	//	loggers.ErrorLogger.Println(err)
	//	return
	//}
	//templates.TemplateCache = cache
	storage.ConnectStorage()

	server := api.APIServer{Addr: ":8000"}
	server.Run()
}
