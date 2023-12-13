document.getElementById('createForm').addEventListener('submit', function(e) {
    e.preventDefault();

    // Создаем новый объект FormData
    var formData = new FormData(this);
    const firstname = formData.get('firstname');
    const lastname = formData.get('lastname');
    const lastname2 = formData.get('lastname2');
    const number = formData.get('number')
    const date_of_birth = formData.get('date_of_birth');
    const place_of_birth = formData.get('birth_place');
    const enrol_date = formData.get('enrol_date');
    const housing_number = formData.get('housing_number');
    const enrol_number = formData.get('enrol_number');
    const housing = formData.get('housing');
    //TODO:profile pic
    //TODO: сделать ввод dd.mm.yyyy, который прогонит в yyyy-mm-dd

    // Создаем объект для отправки данных в формате JSON
    var data = {
        full_name: (firstname + " " + lastname + " " + lastname2),
        card_number: number,
        birth_date: date_of_birth,
        housing_order_number: housing_number,
        enrollment_order_number: enrol_number,
        enrollment_date: enrol_date,
        birth_place: place_of_birth,
        residence_address: housing,

    };

    fetch("/api/addstudent", {
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
                document.body.append("ВСЕ СОЗДАНО")
            }
            if (response.status == 400){
                document.body.append("ВЫ НЕПРАВИЛЬНО ВВЕЛИ ДАННЫЕ")
            }
        })
        .catch(error => console.error('Error:', error));
})
function convertDate(dateString) {
    const parts = dateString.split('.');
    return parts[2] + '-' + parts[1] + '-' + parts[0];
}

