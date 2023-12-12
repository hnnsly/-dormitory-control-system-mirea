package routes

import (
	"github.com/gin-gonic/gin"
	"hackaton/internal/controllers/loginPage"
	"hackaton/internal/controllers/studentCards"
	"hackaton/internal/controllers/studentList"
)

func Setup(app *gin.Engine) {
	app.LoadHTMLGlob("web/html/*")
	app.Static("/static", "./web/static")

	app.GET("/login", loginPage.Start)
	app.GET("/students/find", studentList.FindPage)
	app.POST("/students/find", studentList.ListStudents)
	app.GET("/students/show", studentCards.ShowStudentCard)
	app.GET("/students/add", studentList.AddStudentPage)
	app.GET("/students/edit", studentList.EditStudentPage) //TODO: из-за такого при первом заходе на поиск будет нихуя, а после поиска обновление страницы
	app.POST("/api/register", loginPage.Register)
	app.POST("/api/login", loginPage.Login)
	app.GET("/api/user", loginPage.User)
	app.POST("/api/addstudent", studentList.AddStudentAPI)
	app.POST("/api/editstudent", studentList.EditStudentAPI)
	app.POST("/api/logout", loginPage.Logout)
	// TODO: Роутер к главной странице, судя по всему и гет и пост
	//app.GET("/login", controllers.Logout)
}
