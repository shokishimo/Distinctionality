var questionNo = 1;
var answerOrQuestion = true; // true == answer, false = question to show on button

var questionNoArea;
var question;
var leftBtn;
var rightBtn;
var midBtn;

window.onload = function() {
    questionNoArea = document.querySelector('.No');
    question = document.getElementById('qanda');
    leftBtn = document.querySelector('.left');
    rightBtn = document.querySelector('.right');
    midBtn = document.querySelector('.flip');

    // question number
    questionNoArea.innerHTML = 'Q' + questionNo;
    // get question from sessionStorage
    var qData = sessionStorage.getItem('Question'+questionNo);
    // Question text
    question.innerHTML = qData;

    midBtn.innerHTML = 'answer';

    if (questionNo == 1) {
        leftBtn.setAttribute("hidden", "hidden");
    }
};

function goBack() {
    questionNo--;
    if (questionNo == 1) {
        leftBtn.setAttribute("hidden", "hidden");
    } else {
        leftBtn.removeAttribute("hidden");
    }

    questionNoArea.innerHTML = 'Q' + questionNo;
    valueToShow = sessionStorage.getItem('Question'+questionNo);
    question.innerHTML = valueToShow;
    midBtn.innerHTML = "answer"
    answerOrQuestion = true;
};

function goNext() {
    if (questionNo == 20) {
        goToEndPage();
    } else {
        questionNo++;
        if (questionNo == 2) {
            leftBtn.removeAttribute("hidden");
        }

        questionNoArea.innerHTML = 'Q' + questionNo;
        valueToShow = sessionStorage.getItem('Question'+questionNo);
        question.innerHTML = valueToShow;
        midBtn.innerHTML = "answer"
        answerOrQuestion = true;
    }
};

function flip() {
    var valueToShow;

    if (answerOrQuestion) {
        valueToShow = sessionStorage.getItem('Answer'+questionNo);
        question.innerHTML = valueToShow;
        midBtn.innerHTML = "question"
        answerOrQuestion = false;
    } else {
        valueToShow = sessionStorage.getItem('Question'+questionNo);
        question.innerHTML = valueToShow;
        midBtn.innerHTML = "answer"
        answerOrQuestion = true;
    }
};

function goToEndPage() {

}