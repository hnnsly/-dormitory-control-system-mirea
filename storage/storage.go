package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"math/rand"
	"strconv"
)

var Store PStorage

func ConnectStorage() {

	Store.createStorage()

	err := Store.InitResidences()
	if err != nil {
		panic(fmt.Sprintf("could not fill 'Общага' table: %v", err))
	}
	err = Store.InitStudents()
	if err != nil {
		panic(fmt.Sprintf("could not fill 'Студенты' table: %v", err))
	}

}

func (store *PStorage) createStorage() {
	connectionString := "host=localhost port=5432 user=postgres password=testtest dbname=dormitory sslmode=disable"
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic("could not connect to the database")
	}

	store.Db = db

	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("could not ping the database: %v", err))
	}

	store.initTables()
}

func (store *PStorage) initTables() {
	_, err := store.Db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL
		);
	`)
	if err != nil {
		panic(fmt.Sprintf("could not create 'users' table: %v", err))
	}

	_, err = store.Db.Exec(`
		CREATE TABLE IF NOT EXISTS students (
		    id SERIAL PRIMARY KEY,
			card_number VARCHAR(255),
    		full_name VARCHAR(255),
    		birth_date VARCHAR(255),
    		photo_url TEXT,
    		housing_order_number VARCHAR(255),
    		enrollment_order_number VARCHAR(255),
    		enrollment_date VARCHAR(255),
    		birth_place VARCHAR(255),
    		residence_address VARCHAR(255),
    		residence_id INT
		);
	`)

	_, err = store.Db.Exec(`CREATE TABLE IF NOT EXISTS residences (
			id SERIAL PRIMARY KEY,
			address VARCHAR(255),
			floor INTEGER,
			room INTEGER,
			place INTEGER,
			is_occupied INTEGER
		)
`)
	if err != nil {
		panic(fmt.Sprintf("could not create 'students' table: %v", err))
	}

}

type PStorage struct { // PStorage - PostgreSQL Store
	Db *sql.DB
}

func (store *PStorage) InitStudents() error {
	var rowCount int
	store.Db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", "students")).Scan(&rowCount)
	if rowCount != 0 {
		return nil
	}

	cardNumber := func() string {
		return strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10))
	}

	fullNames := []string{
		"Иванов Иван Иванович", "Петров Петр Петрович", "Сидоров Сидор Сидорович",
		"Козлов Козел Козлович", "Михайлов Михаил Михайлович", "Кузнецов Кузнец Кузнецович",
		"Попов Поп Попович", "Смирнов Смирн Смирнович", "Козлов Козл Козлович",
		"Морозов Мороз Морозович", "Васильев Василий Васильевич", "Соколов Сокол Соколович",
		"Лебедев Лебедь Лебедевич", "Кузьмин Кузьма Кузьминич", "Новиков Новик Новикович",
		"Васильев Василий Васильевич", "Смирнов Смирн Смирнович", "Соколов Сокол Соколович",
		"Петров Петр Петрович", "Иванов Иван Иванович", "Морозов Мороз Морозович",
		"Кузнецов Кузнец Кузнецович", "Сидоров Сидор Сидорович", "Новиков Новик Новикович",
		"Соколов Сокол Соколович", "Лебедев Лебедь Лебедевич", "Козлов Козел Козлович",
		"Попов Поп Попович", "Михайлов Михаил Михайлович", "Васильев Василий Васильевич",
		"Кузьмин Кузьма Кузьминич", "Смирнов Смирн Смирнович", "Морозов Мороз Морозович",
		"Соколов Сокол Соколович", "Лебедев Лебедь Лебедевич", "Новиков Новик Новикович",
		"Петров Петр Петрович", "Иванов Иван Иванович", "Сидоров Сидор Сидорович",
		"Козлов Козел Козлович", "Михайлов Михаил Михайлович", "Кузнецов Кузнец Кузнецович",
		"Попов Поп Попович", "Смирнов Смирн Смирнович", "Козлов Козл Козлович",
		"Морозов Мороз Морозович", "Васильев Василий Васильевич", "Соколов Сокол Соколович",
		"Лебедев Лебедь Лебедевич", "Кузьмин Кузьма Кузьминич", "Новиков Новик Новикович",
		"Васильев Василий Васильевич", "Смирнов Смирн Смирнович", "Соколов Сокол Соколович",
		"Петров Петр Петрович", "Иванов Иван Иванович", "Морозов Мороз Морозович",
		"Кузнецов Кузнец Кузнецович", "Сидоров Сидор Сидорович", "Новиков Новик Новикович",
		"Соколов Сокол Соколович", "Лебедев Лебедь Лебедевич", "Козлов Козел Козлович",
		"Попов Поп Попович", "Михайлов Михаил Михайлович", "Васильев Василий Васильевич",
		"Кузьмин Кузьма Кузьминич", "Смирнов Смирн Смирнович", "Морозов Мороз Морозович",
		"Соколов Сокол Соколович", "Лебедев Лебедь Лебедевич", "Новиков Новик Новикович",
		"Петров Петр Петрович", "Иванов Иван Иванович", "Сидоров Сидор Сидорович",
		"Козлов Козел Козлович", "Михайлов Михаил Михайлович", "Кузнецов Кузнец Кузнецович",
		"Попов Поп Попович", "Смирнов Смирн Смирнович", "Козлов Козл Козлович",
		"Морозов Мороз Морозович", "Васильев Василий Васильевич", "Соколов Сокол Соколович",
		"Лебедев Лебедь Лебедевич", "Кузьмин Кузьма Кузьминич", "Новиков Новик Новикович",
		"Васильев Василий Васильевич", "Смирнов Смирн Смирнович", "Соколов Сокол Соколович",
		"Петров Петр Петрович", "Иванов Иван Иванович", "Морозов Мороз Морозович",
		"Кузнецов Кузнец Кузнецович", "Сидоров Сидор Сидорович", "Новиков Новик Новикович",
		"Соколов Сокол Соколович", "Лебедев Лебедь Лебедевич", "Козлов Козел Козлович"}
	photos := []string{"https://thumb.tildacdn.com/tild6439-3438-4437-b065-303734623661/-/resize/880x/-/format/webp/_1_5.png", "https://www.ixbt.com/img/n1/news/2023/3/3/ixbtmedia_a_student_cheats_on_an_exam_and_is_worried_not_to_be__94f1ef9f-e6fd-4f63-85d9-90c6a351c619_large.png", "https://graziamagazine.ru/upload/attach/e6b/e6b512b4232be912a087f9166040e6ec.jpg", "https://mcmag.ru/wp-content/uploads/2022/12/stil-v-filme-amerikanskij-psihopat-02-mcmag.ru_.png"}
	birthDates := func() string { return fmt.Sprintf("%v.%v.200%v", rand.Intn(30), rand.Intn(13), rand.Intn(10)) }
	photoURLs := func() string { return photos[rand.Intn(4)] }
	housingOrderNumbers := func() string { return fmt.Sprintf("%v-%v%v", rand.Intn(100), rand.Intn(100), rand.Intn(60)) }
	enrollmentOrderNumbers := func() string { return fmt.Sprintf("%v-%v%v", rand.Intn(100), rand.Intn(100), rand.Intn(60)) }
	enrollmentDates := func() string { return fmt.Sprintf("%v.%v.20%v", rand.Intn(30), rand.Intn(13), rand.Intn(6)+16) }
	birthPlaces := func() string { return fmt.Sprintf("Город%v%v", rand.Intn(10), rand.Intn(100)) }

	for i := 0; i < len(fullNames); i++ {
		temp := Student{
			CardNumber:            cardNumber(),
			FullName:              fullNames[i],
			BirthDate:             birthDates(),
			PhotoUrl:              photoURLs(),
			HousingOrderNumber:    housingOrderNumbers(),
			EnrollmentOrderNumber: enrollmentOrderNumbers(),
			EnrollmentDate:        enrollmentDates(),
			BirthPlace:            birthPlaces(),
			ResidenceAddress:      "",
			ResidenceID:           0,
		}
		store.Add(&temp)
	}

	return nil
}

func (store *PStorage) InitResidences() error {
	var rowCount int
	store.Db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", "residences")).Scan(&rowCount)
	if rowCount != 0 {
		return nil
	}
	addresses := []string{"Ул. Пушкина, д. 21", "Ул. Вернадского, д.86к4", "Ул. Асанова, д.14к8"}
	c := 0
	for _, address := range addresses {
		for floor := 1; floor <= rand.Intn(3)+4; floor++ {
			for room := 1; room <= rand.Intn(120)+20; room++ {
				lim := 0
				if room%2 == 0 {
					lim = 3
				} else {
					lim = 2
				}
				for place := 1; place <= lim; place++ {
					c++
					_, err := store.Db.Exec("INSERT INTO residences VALUES ($1, $2, $3, $4, $5,$6)", c, address, floor, room, place, 0)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}
