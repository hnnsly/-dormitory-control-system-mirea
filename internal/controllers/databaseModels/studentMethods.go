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

// ShowStudentsByCriteria получает на вход название столбца и нужное для отбора значение,
// возвращает []string с ФИО всех студентов, подходящих под критерии
func (m *StudentModel) ShowStudentsByCriteria(column, value string, offset int) ([][]Student, error) {

	query := fmt.Sprintf("SELECT * FROM students WHERE %s = $1 OFFSET $2 LIMIT $3", column)

	rows, err := m.DB.Query(query, value, offset, offset+20)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return nil, fmt.Errorf("Ошибка выполнения запроса")
	}
	defer rows.Close()

	students := make([][]Student, 0, 4)
	studentSMOL := make([]Student, 0, 4)
	var count int
	var block int
	for rows.Next() {
		var user Student
		err := rows.Scan(
			&user.ID,
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
		studentSMOL = append(studentSMOL, user)
		count++
		if count%4 == 0 {
			students = append(students, studentSMOL)
			block++
			count = 0
			studentSMOL = make([]Student, 0, 4)
		}
	}
	if count%4 != 0 {
		students = append(students, studentSMOL)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Студент не найден")
		}
		return nil, fmt.Errorf("Ошибка выполнения запроса")
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

func (m *StudentModel) Rewrite(student Student) error {
	query := `
		UPDATE students
		SET
			full_name = $2,
			birth_date = $3,
			photo_url = $4,
			housing_order_number = $5,
			enrollment_order_number = $6,
			enrollment_date = $7,
			birth_place = $8,
			residence_address = $9
		WHERE
			id = $1
	`

	_, err := StudentsDB.DB.Exec(
		query,
		student.ID,
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
