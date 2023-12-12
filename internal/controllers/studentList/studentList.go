package studentList

import (
	"github.com/gin-gonic/gin"
	"hackaton/internal/controllers/databaseModels"
	"hackaton/pkg/loggers"
)

func ListStudents(c *gin.Context) {
	//var filters map[string]string
	//
	//if err := c.BindJSON(&filters); err != nil {
	//	c.JSON(400, gin.H{"message": "invalid request"})
	//	return
	//}

	templateData, err := databaseModels.StudentsDB.Ping("card_number", "228")
	if err != nil {
		loggers.ErrorLogger.Println(err)
		c.JSON(500, gin.H{"message": "database error"})
		return
	}

	c.HTML(200, "students.html", gin.H{
		"title":    "Login",
		"Students": templateData,
	})
}

// TODO: Если честно не понял что это
