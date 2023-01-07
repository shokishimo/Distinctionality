#### Go web application where you can take English words quizzes.
### Here is the URL to access the site. -> https://distinctionality.onrender.com

##### Project Structure

```
.
|---db
|   |---db.go
|
|---internal
|   |---handler
|   |   |---home.go
|   |   |---insertData.go
|   |   |---quiz.go
|   |   |---read20Quizzes.go
|   |---server.go
|
|---model
|   |---model.go
|
|---static
|   |---css
|   |---js
|   |---template
|
|---main.go
```

##### API Endpoint
/get20Quizzes - to get 20 random quiz data from database. HTTP GET. A request needs 'vsersion' parameter (= '1' or '2') and 'level' parameter (= '1', '2', '3' or '4') to retrieve data. Response body has 20 random quiz data in JSON.
