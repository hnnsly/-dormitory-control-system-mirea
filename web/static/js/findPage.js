document.getElementById('findForm').addEventListener('submit', function(e) {
    e.preventDefault();

    // Создаем новый объект FormData
    var formData = new FormData(this);
    const name = formData.get('name');
    const number = formData.get('number');
    var housing = formData.get('housing');

    // Создаем объект для отправки данных в формате JSON
    var data = {
        name: name,
        number: number,
        housing: housing
    };

    fetch("/students/find", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(response => {
            if (response.redirected) {
                window.location.href = response.url;
            } else {
                return response.text(); // Получаем HTML-страницу в виде текста
            }
        })
        .then(html => {
            document.open(); // Открываем новый документ
            document.write(html); // Пишем HTML-страницу в документ
            document.close(); // Закрываем документ
        })
        .then(data => {
            console.log(data);
            // Дополнительные действия после отправки данных, если необходимо
        })
        .catch(error => console.error('Error:', error));
})


