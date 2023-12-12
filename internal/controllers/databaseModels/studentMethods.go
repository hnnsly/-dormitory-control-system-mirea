package databaseModels

import (
	"database/sql"
	"fmt"
	"hackaton/pkg/database"
)

type StudentModel struct {
	DB *sql.DB
}

var StudentsDB StudentModel

func init() {
	StudentsDB.DB = database.DB
}

func (m StudentModel) Ping(column, value any) ([]Student, error) {
	// Строка запроса
	query := fmt.Sprintf("SELECT * FROM students WHERE %s = ?", column)

	// Выполняем запрос
	rows, err := StudentsDB.DB.Query(query, column)
	if err != nil {
		return nil, fmt.Errorf("Ошибка выполнения запроса")
	}
	defer rows.Close()

	// Срез для хранения результатов
	var students []Student

	// Обработка результатов запроса
	for rows.Next() {
		var user Student
		err := rows.Scan(
			&user.CardNumber,
			&user.FullName,
			&user.BirthDate,
			&user.PhotoUrl,
			&user.HousingOrderNumber,
			&user.EnrollmentOrderNumber,
			&user.EnrollmentDate,
			&user.BirthPlace,
			&user.ResidenceAddress,
		)
		if err != nil {
			return nil, fmt.Errorf("Ошибка обработки результатов запроса")
		}
		// Добавляем пользователя в срез
		students = append(students, user)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Студент не найден")
		}
		return nil, fmt.Errorf("Ошибка выполнения запроса")
	}

	return students, nil
}
