package login

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"hackaton/log"
	"hackaton/storage"
)

func Start(c *gin.Context) {
	c.HTML(200, "index.page.tmpl.html", nil)
}

func Search(c *gin.Context) {
	filter := map[string]string{}
	fmt.Println(c.Request.Body)
	err := c.BindJSON(&filter)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Status(500)
		return
	}
	fmt.Println(filter)
	student, err := storage.Store.ShowStudentsByCriteria(filter["option"], filter["query"], 0)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Status(500)
		return
	}
	fmt.Println(student)
	studJSON, err := json.Marshal(student)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Status(500)
		return
	}
	fmt.Println(string(studJSON))
	c.JSON(200, studJSON)
}
