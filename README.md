# **Student Records API**

A simple CRUD (Create, Read, Update, Delete) API for managing student records, built using **Go** and **SQLite**.

---

## **Project Structure**

```
.
├── config
│   └── local.yml        # Configuration file
├── cmd
│   └── api
│       └── main.go      # Main API server entry point
├── db
│   └── development.db   # SQLite database file
└── ...
```

---

## **Configuration**

The configuration file is located at `config/local.yml` and contains the following settings:

```yaml
env: "dev"                       # Environment (dev, test, prod)
db_path: "db/development.db"     # Path to SQLite database
http_server:
  address: "localhost:8080"      # HTTP server address
```

---

## **Getting Started**

### **Install Dependencies**
Run the following command to install all required Go modules:

```bash
go mod tidy
```

### **Run the API Server**
To start the API server:

```bash
go run cmd/api/main.go
```

---

## **API Endpoints**

### **1. Create a Student**
- **URL**: `/api/students`  
- **Method**: `POST`  
- **Request Body**:

  ```json
  {
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 20
  }
  ```

- **Response**:

  ```json
  {
    "success": 200,
    "id": 1
  }
  ```

---

### **2. Get a Student by ID**
- **URL**: `/api/students/{id}`  
- **Method**: `GET`  
- **Response**:

  ```json
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 20
  }
  ```

---

### **3. Get All Students**
- **URL**: `/api/students`  
- **Method**: `GET`  
- **Response**:

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

---

## **Dependencies**

- [**github.com/ilyakaznacheev/cleanenv**](https://github.com/ilyakaznacheev/cleanenv)  
  Used for loading configuration from YAML files.

- [**modernc.org/sqlite**](https://modernc.org/sqlite)  
  SQLite driver for Go.

- [**github.com/go-playground/validator/v10**](https://github.com/go-playground/validator)  
  Used for validating request payloads.

---

## **License**

This project is licensed under the **MIT License**.  
See the [LICENSE](LICENSE) file for details.
