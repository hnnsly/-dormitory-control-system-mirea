package databaseModels

import (
	"database/sql"
	"fmt"
	"hackaton/pkg/database"
	"hackaton/pkg/loggers"
)

type StudentModel struct {
	DB *sql.DB
}

var StudentsDB StudentModel

func InitStudentsDB() {
	StudentsDB.DB = database.DB
}

// Ping получает на вход название столбца и нужное для отбора значение,
// возвращает []string с ФИО всех студентов, подходящих под критерии
func (m *StudentModel) Ping(column, value any) ([]Student, error) {

	query := fmt.Sprintf("SELECT * FROM students WHERE %s = $1", column)

	rows, err := m.DB.Query(query, value)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return nil, fmt.Errorf("Ошибка выполнения запроса")
	}
	defer rows.Close()

	var students []Student

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
			loggers.ErrorLogger.Println(err)
			return nil, fmt.Errorf("Ошибка обработки результатов запроса")
		}

		students = append(students, user)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Студент не найден")
		}
		return nil, fmt.Errorf("Ошибка выполнения запроса")
	}

	fullNames := make([]string, len(students))

	for i := 0; i < len(students); i++ {
		fullNames[i] = students[i].FullName
	}

	return students, nil
}

func (m *StudentModel) Add(student *Student) error {

	query := `
		INSERT INTO students (
			card_number,
			full_name,
			birth_date,
			photo_url,
			housing_order_number,
			enrollment_order_number,
			enrollment_date,
			birth_place,
			residence_address
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := StudentsDB.DB.Exec(
		query,
		student.CardNumber,
		student.FullName,
		student.BirthDate,
		student.PhotoUrl,
		student.HousingOrderNumber,
		student.EnrollmentOrderNumber,
		student.EnrollmentDate,
		student.BirthPlace,
		student.ResidenceAddress,
	)

	return err
}
