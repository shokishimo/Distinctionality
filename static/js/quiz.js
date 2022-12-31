var questionNo = 1;

var questionNoArea = document.querySelector('.No');
var question = document.getElementById('qanda');

window.onload = function() {
    // question number
    questionNoArea.innerHTML = 'Q' + questionNo;

    // get question from sessionStorage
    var qData = sessionStorage.getItem('Question'+questionNo);

    // Question text
    question.innerHTML = qData;
};