package main

import (
	"github.com/gin-gonic/gin"
	"hackaton/pkg/database"
	"hackaton/routes"
)

func main() {
	database.Connect()

	app := gin.Default()

	routes.Setup(app)

	app.Run(":8000")
}
