# Yet another proposal for a hexagonal architecture in Go

**This is a proposal for implementing a hexagonal architecture in Go using the most popular frameworks and libs.**

The primary goal of this project is to demonstrate how to structure a Go application with a hexagonal architecture using Fiber.

This project is based on and forked from [go-gin-clean-arch](https://github.com/josepdcs/go-proposal-hexagonal-arch)

In this project, I have aimed to adhere to Golang naming conventions and best practices, including:
- [Effective Go](https://golang.org/doc/effective_go.html)
  - https://go.dev/doc/effective_go#package-names
  - https://go.dev/blog/package-names

## Template Structure

- [Fiber](https://gofiber.io/) is a Go web framework built on top of Fasthttp, the fastest HTTP engine for Go. It's designed to ease things up for fast development with zero memory allocation and performance in mind.
- [JWT](github.com/golang-jwt/jwt) A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens.
- [GORM](https://gorm.io/index.html) with [PostgresSQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)The fantastic ORM library for Golang aims to be developer friendly.
- [Wire](https://github.com/google/wire) is a code generation tool that automates connecting components using dependency injection.
- [Koanf](https://github.com/knadh/koanf) is a library for reading configuration from different sources in different formats in Go applications. It is a cleaner, lighter [alternative to spf13/viper](https://github.com/knadh/koanf#alternative-to-viper) with better abstractions and extensibility and far fewer dependencies..
- [Swag](https://github.com/swaggo/swag) converts Go annotations to Swagger Documentation 2.0 with [fiber-swagger](https://github.com/gofiber/swagger) and [swaggo files](github.com/swaggo/files)
- [Testify](https://github.com/stretchr/testify) is a set of packages that provide many tools for testifying that your code will behave as you intend. Features include: easy assertions, mocking, testing suite interfaces and functions, and a test runner.

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

