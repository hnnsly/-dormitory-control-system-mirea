package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hackaton/internal/controllers/databaseModels"
	"hackaton/internal/routes"
	"hackaton/pkg/database"
	"hackaton/pkg/loggers"
	"hackaton/pkg/templates"
)

func main() {
	database.Connect()
	databaseModels.InitStudentsDB()
	cache, err := templates.NewTemplateCache("web/html/")
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return
	}
	templates.TemplateCache = cache
	app := gin.Default()
	//student := &databaseModels.Student{
	//	CardNumber:            333,
	//	FullName:              "lox lox111asddsa",
	//	BirthDate:             time.Now(),
	//	PhotoUrl:              "data:image/jpeg;base64,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",
	//	HousingOrderNumber:    3331,
	//	EnrollmentDate:        time.Now(),
	//	EnrollmentOrderNumber: 234234324,
	//	BirthPlace:            "magadan",
	//	ResidenceAddress:      "norilsk"}
	//err = databaseModels.StudentsDB.Add(student)
	//if err != nil {
	//	loggers.ErrorLogger.Fatal(err)
	//}
	fmt.Println(3 / 2)
	routes.Setup(app)

	app.Run(":8000")
}
