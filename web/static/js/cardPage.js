// Получаем ссылку на кнопку
const backButton = document.getElementById('backButton');

// Добавляем обработчик события для клика по кнопке
backButton.addEventListener('click', function() {
    // Вызываем метод back() объекта history
    history.back();
});