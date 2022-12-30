
// Get a reference to the button element
var button = document.getElementById("getData");


// Set the calback func to be executed when the button is clicked
function getData() {
    var level = button.value;
    var requestStr = '/get20Quizzes?version=1&level=' + level

    fetch(requestStr, {method: 'GET'})
        .then(response => response.json())
        .then(data => {
            var count = 1;
            data.forEach(element => {
                sessionStorage.setItem('Question'+count, element.Question);
                sessionStorage.setItem('Answer'+count, element.answer);
                count++;
            });

            fetch(requestStr, {method: 'GET'})
        });
};