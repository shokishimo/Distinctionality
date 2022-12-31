// Set the calback func to be executed when the button is clicked
function getData(param1, param2) {
    var requestStr = '/get20Quizzes?version=' + param1 + '&level=' + param2;

    fetch(requestStr, {method: 'GET'})
        .then(response => response.json())
        .then(data => {
            var count = 1;
            data.forEach(element => {
                sessionStorage.setItem('Question'+count, element.Question);
                sessionStorage.setItem('Answer'+count, element.Answer);
                count++;
            });
        })
        .then(goToQuiz);
};

function goToQuiz() {
    var form = document.createElement('form');
    form.method = 'GET';
    form.action = '/quiz';
    document.body.appendChild(form);
    form.submit();
};