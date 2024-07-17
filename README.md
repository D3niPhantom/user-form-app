# User Management System

This project is a User Management System with a Go backend using Echo framework and an Angular frontend.

## Prerequisites

Before you begin, ensure you have the following installed on your system:

1. Go (version 1.16 or later)
2. Node.js (version 14 or later) and npm
3. PostgreSQL (version 13 or later)
4. Angular CLI (version 15 or later)

## Backend Setup

1. Clone the repository:
   `git clone https://github.com/your-username/your-repo-name.git`
   `cd your-repo-name`

2. Install Go dependencies:
   `go mod tidy`

3. Set up your PostgreSQL database and update the connection string in `cmd/main.go`:

```go
dbConnectionString := "postgres://your-username:your-password@localhost:5432/your-database-name?sslmode=disable"
```

4. Run the backend server:
   `go run cmd/main.go`
   The server should start on http://localhost:8080.

## Frontend Setup

1. In another terminal, navigate to the frontend directory:
   `cd frontend`
2. Install npm dependencies:
   `npm install`
3. Start the Angular development server:
   `ng serve`
   The frontend should be accessible at http://localhost:4200.

### Additional Dependencies

The project uses the following main dependencies:

#### Backend (Go)

    Echo framework: github.com/labstack/echo/v4
    PostgreSQL driver: github.com/lib/pq
    Squirrel (SQL query builder): github.com/Masterminds/squirrel

#### Frontend (Angular)

    Angular Material
    RxJS

### API Documentation

API documentation is available via Swagger UI. After starting the backend server, visit:
http://localhost:8080/swagger/index.html

## Running Tests

### Backend Tests

In the backend directory, execute `go test ./...` from your terminal.

### Frontend Tests

In the frontend directory, execute `ng test` from your terminal.
