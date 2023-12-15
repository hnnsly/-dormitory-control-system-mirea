package main

import (
	"fmt"
	"hackaton/api"
	"hackaton/storage"
	"hackaton/utils"
)

func main() {

	storage.ConnectStorage()
	utils.TemplateCache, _ = utils.NewTemplateCache("web/html/")
	fmt.Println(utils.TemplateCache)

	server := api.APIServer{Addr: ":8000"}
	server.Run()
}
