package main

import (
	"hackaton/api"
	"hackaton/storage"
)

func main() {

	storage.ConnectStorage()

	student := &storage.Student{
		FullName:              "Игорь Петрович Тортуга",
		CardNumber:            "1448",
		BirthDate:             "01.01.1788",
		PhotoUrl:              "https://i0.wp.com/rollercoasteryears.com/wp-content/uploads/Thrive-During-Finals-.jpg?fit=1000%2C667&ssl=1",
		HousingOrderNumber:    "229",
		EnrollmentDate:        "01.03.2023",
		EnrollmentOrderNumber: "1338",
		ResidenceAddress:      "1",
		BirthPlace:            "Магадан",
	}

	storage.Store.AddStudent(student)

	server := api.APIServer{Addr: ":8000"}
	server.Run()
}
