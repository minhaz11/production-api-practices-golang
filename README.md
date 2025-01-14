# Sample APIs 

This project is a simple CRUD (Create, Read, Update, Delete) API for managing student records. It is built using Go and SQLite.

## Project Structure

## Configuration

The configuration file is located at `config/local.yml`. It contains the following settings:

```yml
env: "dev"
db_path: "db/development.db"
http_server:
  address: "localhost:8080"
```

## Running the Project

Install dependencies:

```
go mod tidy

```

Run the API server:

```
go run cmd/api/main.go 

```




## Sample Student Record

API Endpoints
Create a Student
-URL: /api/students
-Method: POST
-Request Body:

```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 20
}
```

Response Body:

```json
{
  "success": 200,
  "id": 1
}
```

Get a Student by ID
-URL: /api/students/{id}
-Method: GET
-Response:

 ```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 20
}
```


Get All Students
URL: /api/students
Method: GET
Response:

```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 20
  },
  {
    "id": 2,
    "name": "Jane Doe",
    "email": "jane.doe@example.com",
    "age": 22
  }
]
```


##Dependencies

github.com/ilyakaznacheev/cleanenv - For loading configuration from a YAML file.

modernc.org/sqlite - SQLite driver for Go.

github.com/go-playground/validator/v10 - For validating request payloads.

##License
This project is licensed under the MIT License.