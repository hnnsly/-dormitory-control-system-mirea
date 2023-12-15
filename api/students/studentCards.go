package students

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hackaton/log"
	"hackaton/storage"
	"hackaton/utils"
	"net/http"
	"strconv"
)

func ShowStudentCard(c *gin.Context) {
	_, err := utils.CheckJWTAuth(c)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Redirect(302, "/login")
		return
	}
	fmt.Println(c.Query("id"))
	stud, err := storage.Store.ShowStudentsByCriteria("id", c.Query("id"), 0)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Status(http.StatusInternalServerError)
	}
	if len(stud) == 0 {
		err = utils.TemplateCache["card.page.tmpl.html"].Execute(c.Writer, nil)
		if err != nil {
			log.ErrorLogger.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
	} else {
		err = utils.TemplateCache["card.page.tmpl.html"].Execute(c.Writer, stud[0][0])
		if err != nil {
			log.ErrorLogger.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
	}

}
func EditStudentPage(c *gin.Context) {
	_, err := utils.CheckJWTAuth(c)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Redirect(302, "/login")
		return

	}
	stud, err := storage.Store.ShowStudentsByCriteria("id", c.Query("id"), 0)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Status(http.StatusInternalServerError)
	}
	if len(stud) == 0 {
		err = utils.TemplateCache["card.page.tmpl.html"].Execute(c.Writer, nil)
		if err != nil {
			log.ErrorLogger.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
	} else {
		err = utils.TemplateCache["edit.page.tmpl.html"].Execute(c.Writer, stud[0][0])
		if err != nil {
			log.ErrorLogger.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
	}
}
func AddStudentPage(c *gin.Context) {
	_, err := utils.CheckJWTAuth(c)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Redirect(302, "/login")
		return

	}

	err = utils.TemplateCache["add.page.tmpl.html"].Execute(c.Writer, nil)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
}

func EditStudentAPI(c *gin.Context) {
	_, err := utils.CheckJWTAuth(c)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Redirect(302, "/login")
		return
	}

	var filter map[string]string

	if err := c.BindJSON(&filter); err != nil {
		log.ErrorLogger.Println(err)
		c.JSON(400, gin.H{"message": "invalid request"})
		return
	}
	id, err := strconv.Atoi(filter["id"])
	if err != nil {
		log.ErrorLogger.Println(err)
		c.JSON(400, gin.H{"message": "invalid request"})
		return
	}
	stud, _ := storage.Store.ShowStudentsByCriteria("id", strconv.Itoa(id), 0)
	if len(stud) == 0 {
		c.Redirect(302, "/students/find")
		return
	}
	student := &storage.Student{
		ID:                    id,
		CardNumber:            filter["card_number"],
		FullName:              filter["full_name"],
		BirthDate:             filter["birth_date"],
		HousingOrderNumber:    filter["housing_order_number"],
		EnrollmentDate:        filter["enrollment_date"],
		EnrollmentOrderNumber: filter["enrollment_order_number"],
		BirthPlace:            filter["birth_place"],
		PhotoUrl:              stud[0][0].PhotoUrl}
	err = storage.Store.Rewrite(*student)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Status(500)
		return
	}
	c.Redirect(302, "/students/show?id="+strconv.Itoa(id))

}
func AddStudentAPI(c *gin.Context) {
	_, err := utils.CheckJWTAuth(c)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Redirect(302, "/login")
		return
	}

	var filter map[string]string

	if err := c.BindJSON(&filter); err != nil {
		fmt.Println("gerre")
		log.ErrorLogger.Println(err)
		c.JSON(400, gin.H{"message": "invalid request"})
		return
	}
	student := &storage.Student{
		CardNumber:            filter["card_number"],
		FullName:              filter["full_name"],
		BirthDate:             filter["birth_date"],
		PhotoUrl:              "https://upload.wikimedia.org/wikipedia/commons/thumb/1/15/Cat_August_2010-4.jpg/2560px-Cat_August_2010-4.jpg",
		HousingOrderNumber:    filter["housing_order_number"],
		EnrollmentDate:        filter["enrollment_date"],
		EnrollmentOrderNumber: filter["enrollment_order_number"],
		BirthPlace:            filter["birth_place"]}
	id, err := storage.Store.Add(student)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Status(500)
		return
	}
	stud, _ := storage.Store.ShowStudentsByCriteria("id", strconv.Itoa(id), 0)
	adr := stud[0][0].ResidenceAddress
	c.JSON(200, gin.H{
		"address": adr,
	})
}

func DeleteAPI(c *gin.Context) {
	_, err := utils.CheckJWTAuth(c)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Redirect(302, "/login")
		return
	}

	var filter map[string]string

	if err := c.BindJSON(&filter); err != nil {
		log.ErrorLogger.Println(err)
		c.JSON(400, gin.H{"message": "invalid request"})
		return
	}
	id, err := strconv.Atoi(filter["id"])
	if err != nil {
		log.ErrorLogger.Println(err)
		c.JSON(400, gin.H{"message": "invalid request"})
		return
	}
	err = storage.Store.Delete(id)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Status(500)
		return
	}
	c.Redirect(302, "/students/show?id="+strconv.Itoa(id))
}
