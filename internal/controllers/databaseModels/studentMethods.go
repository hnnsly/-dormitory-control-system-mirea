package databaseModels

import (
	"database/sql"
	"fmt"
	"hackaton/pkg/database"
	"hackaton/pkg/loggers"
	"log"
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
	fmt.Println(offset)
	query := fmt.Sprintf("SELECT * FROM students WHERE %s = $1 OFFSET $2 LIMIT 12", column)

	rows, err := m.DB.Query(query, value, offset)
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
			birth_place        
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	var studentID int
	err := m.DB.QueryRow(
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
	).Scan(&studentID)

	addr, addrID, err := m.Settle(studentID)

	query = `UPDATE residences SET residence_address = $1, residence_id = $2 WHERE id = $3`

	m.DB.Exec(query, addr, addrID, studentID)

	if err != nil {
		return err
	}

	return nil
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
			residence_address = $9,
			residence_id = $10
		WHERE
			id = $1
	`

	newAddr, addrID, err := StudentsDB.Settle(student.ID)

	_, err = StudentsDB.DB.Exec(
		query,
		student.ID,
		student.FullName,
		student.BirthDate,
		student.PhotoUrl,
		student.HousingOrderNumber,
		student.EnrollmentOrderNumber,
		student.EnrollmentDate,
		student.BirthPlace,
		newAddr,
		addrID,
	)

	return err
}

func (m *StudentModel) Settle(studentID int) (string, int, error) {

	var address string
	var roomID, floor, room int
	err := m.DB.QueryRow("SELECT id, address, floor, room FROM residences WHERE is_occupied = 0 LIMIT 1").Scan(&roomID, &address, &floor, &room)
	if err != nil {
		log.Println("Error querying room:", err)
		return "", 0, err
	}

	_, err = m.DB.Exec("UPDATE residences SET is_occupied = $1 WHERE id = $2", studentID, roomID)
	if err != nil {
		return "", 0, err
	}

	roomInfo := fmt.Sprintf("%s, %d этаж, %d комната", address, floor, room)

	return roomInfo, roomID, nil
}
