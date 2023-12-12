package databaseModels

import (
	"database/sql"
	"hackaton/pkg/database"
)

type StudentModel struct {
	DB *sql.DB
}

var Students StudentModel

func init() {
	Students.DB = database.DB
}

func (m StudentModel) Ping() {

}
