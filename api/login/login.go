package login

import (
	"github.com/gin-gonic/gin"
	"hackaton/log"
	"hackaton/utils"
	"net/http"
)

func Start(c *gin.Context) {
	err := utils.TemplateCache["login.page.tmpl.html"].Execute(c.Writer, nil)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
}
func Redirect(c *gin.Context) {
	c.Redirect(301, "/students/find")
}
