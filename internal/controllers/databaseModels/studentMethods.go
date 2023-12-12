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
	err := m.DB.QueryRow("SELECT id, username, password FROM users WHERE username = $1", data["email"]).
		Scan(&user.Id, &user.Email, &user.Password)

	m.DB.QueryRow("SELECT * FROM users").Scan()
}
