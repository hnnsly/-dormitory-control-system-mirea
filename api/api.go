package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hackaton/api/login"
	"time"
)

func (s *APIServer) Run() {

	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5173"
		},
		MaxAge: 12 * time.Hour,
	}))

	app.GET("/login", login.Start)
	app.POST("/api/register", login.Register)
	app.POST("/api/login", login.Login)
	app.GET("/api/user", login.User)
	app.POST("/api/search", login.Search)
	app.POST("/api/logout", login.Logout)
	// TODO: Роутер к главной странице, судя по всему и гет и пост
	//app.GET("/login", controllers.Logout)

	app.Run(s.Addr)
}

type APIServer struct {
	Addr string
}
