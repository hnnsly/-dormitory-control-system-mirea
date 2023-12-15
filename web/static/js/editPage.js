document.getElementById('editForm').addEventListener('submit', function(e) {
    e.preventDefault();

    // Создаем новый объект FormData
    var formData = new FormData(this);
    const full_name = formData.get('full_name');
    const number = formData.get('number')
    const date_of_birth = formData.get('date_of_birth');
    const place_of_birth = formData.get('birth_place');
    const enrol_date = formData.get('enrol_date');
    const housing_number = formData.get('housing_number');
    const enrol_number = formData.get('enrol_number');
    //TODO:profile pic

    const queryString = window.location.search;

// Создаем объект URLSearchParams, передавая строку запроса
    const params = new URLSearchParams(queryString);

// Получаем значения параметров
    const id = params.get('id');
    // Создаем объект для отправки данных в формате JSON
    var data = {
        id: id,
        full_name: full_name,
        card_number: number,
        birth_date: date_of_birth,
        housing_order_number: housing_number,
        enrollment_order_number: enrol_number,
        enrollment_date: enrol_date,
        birth_place: place_of_birth,

    };

    fetch("/api/editstudent", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(response => {
            if (response.redirected) {
                window.location.href = response.url
            }
            if (response.ok) {
                document.getElementById("statusTab").textContent = "Данные изменены"
            }
            if (response.status == 400){
                document.getElementById("statusTab").textContent = "Вы неправильно ввели данные!"
            }
        })
        .catch(error => console.error('Error:', error));
})

document.getElementById('deleteButton').addEventListener('click', function(e) {
    const userConfirmation = confirm("Вы уверены, что хотите удалить запись студента?");
    if (!userConfirmation) {
        event.preventDefault();
    }

    const queryString = window.location.search;

// Создаем объект URLSearchParams, передавая строку запроса
    const params = new URLSearchParams(queryString);

// Получаем значения параметров
    const id = params.get('id');
    // Создаем объект для отправки данных в формате JSON
    var data = {
        id: id,

    };

    fetch("/api/deletestudent", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(response => {
            if (response.redirected) {
                window.location.href = response.url
            }
            if (response.ok) {
                document.getElementById("statusTab").textContent = "Данные изменены"
            }
            if (response.status == 400){
                document.getElementById("statusTab").textContent = "Вы неправильно ввели данные!"
            }
            if (response.status == 500){
                document.getElementById("statusTab").textContent = "На данный момент все комнаты заняты, попробуйте позже"
            }
        })
        .catch(error => console.error('Error:', error));
})



