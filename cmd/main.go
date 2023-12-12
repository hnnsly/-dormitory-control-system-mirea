package main

import (
	"github.com/gin-gonic/gin"
	"hackaton/internal/routes"
	"hackaton/pkg/database"
)

func main() {
	database.Connect()

	app := gin.Default()

	routes.Setup(app)

	app.Run(":8000")
}
