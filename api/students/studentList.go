package students

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hackaton/log"
	"hackaton/storage"
	"hackaton/utils"
	"net/http"
	"net/url"
	"strconv"
)

func ListStudents(c *gin.Context) {
	fmt.Println(c.Request.URL, "akgjkad")
	_, err := utils.CheckJWTAuth(c)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Redirect(302, "/login")
		return

	}
	encodedQuery := c.Request.URL.RawQuery
	decodedQuery, err := url.QueryUnescape(encodedQuery)
	if err != nil {
		log.ErrorLogger.Println(err)
	}
	queryParams, err := url.ParseQuery(decodedQuery)
	if err != nil {
		log.ErrorLogger.Println(err)
	}
	filters := make(map[string]string)
	for key, values := range queryParams {
		if len(values) > 0 {
			// Берем только первое значение для каждого ключа
			filters[key] = values[0]
		}
	}
	fmt.Println(filters)
	var page int
	if _, ok := filters["page"]; !ok {
		page = 1
	} else {
		page, _ = strconv.Atoi(filters["page"])
	}
	if page == 0 {
		page = 1
	}
	var templateData [][]storage.Student
	fmt.Println((page - 1) * 12)
	if filters["name"] != "" {
		tempData, err := storage.Store.ShowStudentsByCriteria("full_name", filters["name"], (page-1)*12)
		if err != nil {
			log.ErrorLogger.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		templateData = append(templateData, tempData...)
	} else if filters["number"] != "" {
		tempData, err := storage.Store.ShowStudentsByCriteria("card_number", filters["number"], (page-1)*12)
		if err != nil {
			log.ErrorLogger.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		templateData = append(templateData, tempData...)
	}
	if filters["housing"] != "0" {
		tempData, err := storage.Store.ShowStudentsByCriteria("residence_address", filters["housing"], (page-1)*12)
		if err != nil {
			log.ErrorLogger.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		templateData = append(templateData, tempData...)
	}
	u := url.URL{}
	q := u.Query()
	for key, value := range filters {
		if key != "page" {
			q.Set(key, value)
		}
	}
	queryString := q.Encode()
	pageNext := page + 1
	pagePrev := page - 1
	if pagePrev == 0 {
		pagePrev = 1
	}
	if page < 3 {
		page = 3
	}
	data := gin.H{
		"title":    "Login",
		"Students": templateData,
		"Name":     filters["name"],
		"Number":   filters["number"],
		"Page1":    page - 2,
		"Page2":    page - 1,
		"Page3":    page,
		"Page4":    page + 1,
		"Page5":    page + 2,
		"PageNext": pageNext,
		"PagePrev": pagePrev,
		"Query":    queryString,
	}
	err = utils.TemplateCache["students.page.tmpl.html"].Execute(c.Writer, data)
	if err != nil {
		log.ErrorLogger.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	//c.HTML(200, "students.page.tmpl.html", data)
}
