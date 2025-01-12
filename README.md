# Yet another proposal for a hexagonal architecture in Go

**This is a proposal for implementing a hexagonal architecture in Go using the most popular frameworks and libs.**

The primary goal of this project is to demonstrate how to structure a Go application with a hexagonal architecture using Gin.

This project is based on and forked from [go-gin-clean-arch](https://github.com/thnkrn/go-gin-clean-arch)

In this project, I have aimed to adhere to Golang naming conventions and best practices, including:
- [Effective Go](https://golang.org/doc/effective_go.html)
  - https://go.dev/doc/effective_go#package-names
  - https://go.dev/blog/package-names

## Template Structure

- [Gin](github.com/gin-gonic/gin) is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
- [JWT](github.com/golang-jwt/jwt) A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens.
- [GORM](https://gorm.io/index.html) with [PostgresSQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)The fantastic ORM library for Golang aims to be developer friendly.
- [Wire](https://github.com/google/wire) is a code generation tool that automates connecting components using dependency injection.
- [Viper](https://github.com/spf13/viper) is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.
- [swag](https://github.com/swaggo/swag) converts Go annotations to Swagger Documentation 2.0 with [gin-swagger](https://github.com/swaggo/gin-swagger) and [swaggo files](github.com/swaggo/files)

## Using `go-proposal-hexagonal-arch` project

To use `go-proposal-hexagonal-arch` project, follow these steps:

```bash
# Navigate into the project
cd ./go-proposal-hexagonal-arch

# Install dependencies
make deps

# Generate wire_gen.go for dependency injection & build the project
# Please make sure you are export the env for GOPATH
make build

# Run the project in Development Mode
make run
```

Additional commands:

```bash
➔ make help
build                          Generate wire_gen.go && Compile the code, build Executable File
run                            Start application
test                           Run tests
test-coverage                  Run tests and generate coverage file
deps                           Install dependencies
deps-cleancache                Clear cache in Go module
wire                           Generate wire_gen.go
swag                           Generate swagger docs
help                           Display this help screen
```

## Available Endpoint

In the project directory, you can call:

### `GET /`

For getting status page

### `POST /login`

For generating a JWT

### `GET /api/users`

For getting all of users

### `GET /api/users/:id`

For getting user by ID

### `POST /api/users`

For creating new user

### `DELETE /api/users/:id`

For removing existing user

### `PUT /api/users/:id`

For updating existing user

