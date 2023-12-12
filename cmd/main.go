package main

import (
	"github.com/gin-gonic/gin"
	"hackaton/internal/controllers/databaseModels"
	"hackaton/internal/routes"
	"hackaton/pkg/database"
)

func main() {
	database.Connect()
	databaseModels.InitStudents()

	app := gin.Default()
	//_, err := database.DB.Exec("INSERT INTO students VALUES ($1, $2, $3, $4, $5, $6)", "228", "PETR PER FFF", "01.02.2012", "https://i.natgeofe.com/n/548467d8-c5f1-4551-9f58-6817a8d2c45e/NationalGeographic_2572187_square.jpg", "1498", "123")
	//if err != nil {
	//	log.Fatal(err)
	//}
	routes.Setup(app)

	app.Run(":8000")
}
