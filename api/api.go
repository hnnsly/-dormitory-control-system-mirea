package api

import (
	"github.com/gin-gonic/gin"
	"hackaton/api/login"
	"hackaton/api/students"
)

func (s *APIServer) Run() {

	app := gin.Default()

	app.LoadHTMLGlob("web/html/*")
	app.Static("/static", "./web/static")

	app.GET("/login", login.Start)
	app.GET("/students/find", students.ListStudents)
	app.GET("/students/show", students.ShowStudentCard)
	app.GET("/students/add", students.AddStudentPage)
	app.GET("/students/edit", students.EditStudentPage) //TODO: из-за такого при первом заходе на поиск будет нихуя, а после поиска обновление страницы
	app.POST("/api/register", login.Register)
	app.POST("/api/login", login.Login)
	app.GET("/api/user", login.User)
	app.POST("/api/addstudent", students.AddStudentAPI)
	app.POST("/api/editstudent", students.EditStudentAPI)
	app.POST("/api/logout", login.Logout)
	// TODO: Роутер к главной странице, судя по всему и гет и пост
	//app.GET("/login", controllers.Logout)

	app.Run(s.Addr)
}

type APIServer struct {
	Addr string
}
