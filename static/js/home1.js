
// Get a reference to the button element
var button = document.getElementById("getData");

var xhr = new XMLHttpRequest();

// Set the calback func to be executed when the button is clicked
function getData() {
    // Set the callback func to be executed when the request is done
    xhr.onload = function() {
        if (xhr.status == 200) 
        {
            // parse json data
            var data = JSON.parse(xhr.responseText);
            console.log(data);
        }
    };

    var level = button.value;
    var requestStr = "http://localhost:8080/get20Quizzes?version=1&level=" + level

    xhr.open("GET", requestStr);

    xhr.send();
};