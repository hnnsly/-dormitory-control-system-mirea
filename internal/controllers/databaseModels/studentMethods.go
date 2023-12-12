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

// Ping получает на вход название столбца и нужное для отбора значение,
// возвращает []string с ФИО всех студентов, подходящих под критерии
func (m StudentModel) Ping(column, value any) ([]string, error) {
	// Строка запроса
	query := fmt.Sprintf("SELECT * FROM students WHERE %s = ?", column)

	// Выполняем запрос
	rows, err := StudentsDB.DB.Query(query, value)
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

	fullNames := make([]string, len(students))

	for i := 0; i < len(students); i++ {
		fullNames[i] = students[i].FullName
	}

	return fullNames, nil
}

// TODO: если честно я это не тестил, но в теории и по бингу все работает
// да и возможности потестить у меня особо не было, полноценная хтмл страница работает через очко, поскольку бинг,
// а через постмен че то проблематично
