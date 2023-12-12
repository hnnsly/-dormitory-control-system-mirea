package studentList

import (
	"github.com/gin-gonic/gin"
	"hackaton/internal/controllers/databaseModels"
	"hackaton/internal/helping"
	"hackaton/pkg/loggers"
	"net/http"
	"strconv"
)

func FindPage(c *gin.Context) {
	_, err := helping.CheckJWTAuth(c)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		c.Redirect(302, "/login")
		return

	}
	//err = templates.TemplateCache["students.page.tmpl.html"].Execute(c.Writer, nil)
	//if err != nil {
	//	loggers.ErrorLogger.Println(err)
	//	return
	//}
	//TODO: временная шняга
	templateData, err := databaseModels.StudentsDB.ShowStudentsByCriteria("birth_place", "magadan", 0)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.HTML(200, "students.page.tmpl.html", gin.H{
		"title":    "Login",
		"Students": templateData,
	})
}
func ListStudents(c *gin.Context) {
	_, err := helping.CheckJWTAuth(c)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		c.Redirect(302, "/login")
		return

	}
	//var filters map[string]string
	//
	//if err := c.BindJSON(&filters); err != nil {
	//	c.JSON(400, gin.H{"message": "invalid request"})
	//	return
	//}
	page, _ := strconv.Atoi(c.Query("page"))
	if c.Query("page") == "" {
		page = 1
	}
	templateData, err := databaseModels.StudentsDB.ShowStudentsByCriteria("birth_place", "magadan", (page - 1*20))
	if err != nil {
		loggers.ErrorLogger.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.HTML(200, "students.page.tmpl.html", gin.H{
		"title":    "Login",
		"Students": templateData,
	})
}
