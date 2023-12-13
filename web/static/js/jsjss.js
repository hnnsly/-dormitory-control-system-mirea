document.getElementById('authForm').addEventListener('submit', function(e) {
    e.preventDefault();

    // Создаем новый объект FormData
    var formData = new FormData(this);
    const action = formData.get('action');
    const username = formData.get('username');
    const url = action === 'login' ? '/api/login' : '/api/register';
    const password = formData.get('password');

    // Создаем объект для отправки данных в формате JSON
    var data = {
        email: username,
        password: password,
    };

    fetch(url, {
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
        })
        .then(data => {
            console.log(data);
            // Дополнительные действия после отправки данных, если необходимо
        })
        .catch(error => console.error('Error:', error));
})

