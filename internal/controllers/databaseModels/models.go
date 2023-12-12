package databaseModels

import "time"

type Student struct {
	CardNumber            int       `json:"card_number"`
	FullName              string    `json:"full_name"`
	BirthDate             time.Time `json:"birth_date"`
	PhotoUrl              string    `json:"photo_url"`
	HousingOrderNumber    int       `json:"housing_order_number"`
	EnrollmentOrderNumber int       `json:"enrollment_order_number"`
	EnrollmentDate        time.Time `json:"enrollment_date"`
	BirthPlace            string    `json:"birth_place"`
	ResidenceAddress      string    `json:"residence_address"`
}
