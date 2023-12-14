package storage

import (
	"database/sql"
	"fmt"
	"hackaton/log"
	"time"
)

// ShowStudentsByCriteria получает на вход название столбца и нужное для отбора значение,
// возвращает []string с ФИО всех студентов, подходящих под критерии
func (st *PStorage) ShowStudentsByCriteria(column, value string, offset int) ([][]Student, error) {
	fmt.Println(offset)
	query := fmt.Sprintf("SELECT * FROM students WHERE %s = $1 OFFSET $2 LIMIT 12", column)

	rows, err := st.Db.Query(query, value, offset)
	if err != nil {
		log.ErrorLogger.Println(err)
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
			log.ErrorLogger.Println(err)
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

func (st *PStorage) SearchStudent(student *Student) error {
	return nil
}

func (st *PStorage) AddStudent(student *Student) error {
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

	_, err := st.Db.Exec(
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

func (st *PStorage) RewriteStudent(student Student) error {
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

	_, err := st.Db.Exec(
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

type Student struct {
	ID                    int       `json:"id"`
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
