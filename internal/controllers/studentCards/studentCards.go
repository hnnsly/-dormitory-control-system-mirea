package studentCards

import (
	"github.com/gin-gonic/gin"
	"hackaton/internal/controllers/databaseModels"
	"hackaton/internal/helping"
	"hackaton/pkg/loggers"
	"hackaton/pkg/templates"
	"net/http"
)

func ShowStudentCard(c *gin.Context) {
	_, err := helping.CheckJWTAuth(c)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		c.Redirect(302, "/login")
		return
	}
	stud, err := databaseModels.StudentsDB.ShowStudentsByCriteria("id", c.Query("id"), 0)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		c.Status(http.StatusInternalServerError)
	}
	err = templates.TemplateCache["students.page.tmpl.html"].Execute(c.Writer, stud[0])
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return
	}
}
