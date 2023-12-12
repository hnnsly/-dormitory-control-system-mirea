package loginPage

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Start(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Login",
	})
}
