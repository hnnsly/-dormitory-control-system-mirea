package login

import (
	"github.com/gin-gonic/gin"
)

func Start(c *gin.Context) {
	c.HTML(200, "index.page.tmpl.html", nil)
}
func Redirect(c *gin.Context) {
	c.Redirect(301, "/students/find")
}
