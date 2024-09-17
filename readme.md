# Golang Boilerplate

This boilerplate provides a starting point for building a web application using the following technologies:

- **Framework**: [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- **ORM**: [gorm.io/gorm](https://gorm.io/gorm)
- **ORM Postgres**: [gorm.io/driver/postgres](https://gorm.io/driver/postgres)
- **Swagger**: [github.com/go-swagger/go-swagger](https://github.com/go-swagger/go-swagger)
- **Logger**: [github.com/uber-go/zap](https://github.com/uber-go/zap)
- **Testing**: [github.com/stretchr/testify](https://github.com/stretchr/testify)
- **Linting**: [github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint)
- **Formatting**: [pkg.go.dev/cmd/gofmt](https://pkg.go.dev/cmd/gofmt)
- **OpenTelemetry**: [opentelemetry.io/docs/languages/go/getting-started/](https://opentelemetry.io/docs/languages/go/getting-started/)

## Getting Started

### Prerequisites

Ensure you have the following installed:

- Go (1.16+)
- PostgreSQL
- Git

### Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/yourusername/golang-boilerplate.git
    cd golang-boilerplate
    ```

2. **Install dependencies:**
    ```sh
    go mod tidy
    ```

3. **Set up environment variables:**
    Create a `.env` file in the root directory and add your configuration:
    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=youruser
    DB_PASSWORD=yourpassword
    DB_NAME=yourdbname
    ```

4. **Run database migrations:**
    ```sh
    go run scripts/migrate.go
    ```

### Running the Application

1. **Start the server:**
    ```sh
    go run main.go
    ```

2. **Access the application:**
    Open your browser and navigate to `http://localhost:8080`

### Generating Swagger Documentation

To generate Swagger documentation, run:

```
    sh
    swag init -g ./modules/router.go -o docs
```

### Testing

Run tests using:

```
sh
go test ./...
```

### Linting

Lint your code using:

```
sh
golangci-lint run
```

### Formatting

Format your code using:

```
sh
gofmt -w .
```


## Contributing

Feel free to submit issues, fork the repository and send pull requests!
