package loginPage

import (
	"github.com/gin-gonic/gin"
	"hackaton/pkg/loggers"
	"hackaton/pkg/templates"
)

func Start(c *gin.Context) {
	err := templates.TemplateCache["index.page.tmpl.html"].Execute(c.Writer, nil)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return
	}
}
