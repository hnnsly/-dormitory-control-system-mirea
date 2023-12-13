package login

import (
	"github.com/gin-gonic/gin"
	"hackaton/utils"
)

func Start(c *gin.Context) {
	err := utils.TemplateCache["index.page.tmpl.html"].Execute(c.Writer, nil)
	if err != nil {
		utils.ErrorLogger.Println(err)
		return
	}
}
