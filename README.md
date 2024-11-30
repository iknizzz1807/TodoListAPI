# Todo List API

## Overview
This is a simple TodoList API developed as a homework project for GDSC UIT (Google Developer Student Clubs, University of Information Technology). The project is a RESTful API built with Go, using the standard `net/http` package and SQLite as the database.

## Features
- Create new todo items
- Retrieve all todo items
- Retrieve a specific todo item by ID
- Retrieve completed todo items
- Update existing todo items
- Delete todo items

## Technology Stack
- Language: Go
- Database: SQLite
- HTTP Framework: Standard library `net/http`

## Endpoints

### Create Todo
- **Route**: `/create`
- **Method**: `POST`
- **Body**: Todo item in JSON (excluding ID)
- **Returns**: Created todo item with assigned ID

### Get All Todos
- **Route**: `/getall`
- **Method**: `GET`
- **Returns**: List of all todo items

### Get Todo by ID
- **Route**: `/getbyid?id=<todo_id>`
- **Method**: `GET`
- **Returns**: Specific todo item matching the ID

### Get Completed Todos
- **Route**: `/getbycompleted`
- **Method**: `GET`
- **Returns**: List of completed todo items

### Update Todo
- **Route**: `/update`
- **Method**: `PUT`
- **Body**: Todo item in JSON (including ID)
- **Returns**: Updated todo item

### Delete Todo
- **Route**: `/delete?id=<todo_id>`
- **Method**: `DELETE`
- **Returns**: No content (204 status code)

## Prerequisites
- Go (1.16+)
- SQLite
- `github.com/mattn/go-sqlite3` package

## Installation
1. Clone the repository
```bash
git clone https://github.com/iknizzz1807/TodoListAPI
cd TodoListAPI
```

2. Install dependencies
```bash
go mod tidy
```

3. Run the application
```bash
go run main.go
```

The server will start on `localhost:8080`

## Todo Item Structure
```json
{
  "id": 1,
  "title": "Complete homework",
  "content": "Finish the TodoList API project",
  "finished": false
}
```

## Concurrency & Safety
- Uses `sync.Mutex` for thread-safe database operations
- Proper error handling for database and HTTP request scenarios

## Future Improvements
- Add authentication
- Implement more advanced filtering
- Create a frontend interface
- Add pagination for todo lists

## Contributing
Feel free to fork the project and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)

## Acknowledgements
Developed as part of the GDSC UIT homework assignment.
