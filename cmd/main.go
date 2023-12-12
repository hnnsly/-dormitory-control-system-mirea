package main

import (
	"github.com/gin-gonic/gin"
	"hackaton/internal/controllers/databaseModels"
	"hackaton/internal/routes"
	"hackaton/pkg/database"
	"hackaton/pkg/loggers"
	"hackaton/pkg/templates"
)

func main() {
	database.Connect()
	databaseModels.InitStudentsDB()
	cache, err := templates.NewTemplateCache("web/html/")
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return
	}
	templates.TemplateCache = cache
	app := gin.Default()
	_, err = database.DB.Exec("INSERT INTO students VALUES ($1, $2, $3, $4, $5, $6)", "2332", "HUY ыфвыффвы FFF", "01.02.2012", "https://i.natgeofe.com/n/548467d8-c5f1-4551-9f58-6817a8d2c45e/NationalGeographic_2572187_square.jpg", "1498", "123")
	if err != nil {
		loggers.ErrorLogger.Fatal(err)
	}
	routes.Setup(app)

	app.Run(":8000")
}
