package routes

import (
	"github.com/gin-gonic/gin"
	"hackaton/internal/controllers"
)

func Setup(app *gin.Engine) {
	app.LoadHTMLGlob("web/html/*")
	app.Static("/static", "./web/static")
	app.GET("/login", controllers.Start)
	app.GET("/map", controllers.Start)
	app.GET("/find", controllers.Start)
	app.GET("/user", controllers.User)
	app.POST("/api/register", controllers.Register)
	app.POST("/api/login", controllers.Login)
	app.GET("/api/user", controllers.User)
	app.POST("/api/logout", controllers.Logout)
	//app.GET("/login", controllers.Logout)
}
