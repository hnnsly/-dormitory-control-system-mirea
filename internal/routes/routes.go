package routes

import (
	"github.com/gin-gonic/gin"
	"hackaton/internal/controllers/loginPage"
	"hackaton/internal/controllers/studentList"
)

func Setup(app *gin.Engine) {
	app.LoadHTMLGlob("web/html/*")
	app.Static("/static", "./web/static")
	app.GET("/login", loginPage.Start)
	app.GET("/find", studentList.ListStudents)
	app.POST("/api/register", loginPage.Register)
	app.POST("/api/login", loginPage.Login)
	app.GET("/api/user", loginPage.User)
	app.POST("/api/logout", loginPage.Logout)
	// TODO: Роутер к главной странице, судя по всему и гет и пост
	//app.GET("/login", controllers.Logout)
}
