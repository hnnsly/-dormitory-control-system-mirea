package api

import (
	"github.com/gin-gonic/gin"
	login2 "hackaton/api/login"
	"hackaton/api/students"
)

func (s *APIServer) Run() {

	app := gin.Default()

	app.Static("/static", "web/static")

	app.GET("/", login2.Redirect)
	app.GET("/login", login2.Start)
	app.GET("/students/find", students.ListStudents)
	app.GET("/students/show", students.ShowStudentCard)
	app.GET("/students/add", students.AddStudentPage)
	app.GET("/students/edit", students.EditStudentPage) //TODO: из-за такого при первом заходе на поиск будет нихуя, а после поиска обновление страницы
	app.POST("/api/register", login2.Register)
	app.POST("/api/login", login2.Login)
	app.GET("/api/user", login2.User)
	app.POST("/api/addstudent", students.AddStudentAPI)
	app.POST("/api/editstudent", students.EditStudentAPI)
	app.POST("/api/deletestudent", students.DeleteAPI)
	app.POST("/api/logout", login2.Logout)
	// TODO: Роутер к главной странице, судя по всему и гет и пост
	//app.GET("/login", controllers.Logout)

	app.Run(s.Addr)
}

type APIServer struct {
	Addr string
}
