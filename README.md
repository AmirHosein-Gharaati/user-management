# Golang user-management service

## Usecases

- Register a new user by email

## Installation

You must have installed golang v1.22

Make sure you have installed the [taskfile](https://taskfile.dev/installation/) tool.

You can use command below to install the required dependencies:

```bash
task install
```

## Usage

You should run database first, and the run the application.

You can run the docker compose by:

```bash
task up
```

And the run the application:

```
task run
```

## Project structure

Here we used hexagonal artchitecture for this project.

The project structure is constructed by this conventions:

- `cmd`: the main point of the program
- `docker`: everything related to the docker files
- `docs`: documentation of the projects like database docs
- `internal`: the application code and implementations

## Learning References

You can read more about hexagonal architecture:

- [Hexagonal Architecture in GO](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3)
- [Hexagonal Architecture, there are always two sides to every story](https://medium.com/ssense-tech/hexagonal-architecture-there-are-always-two-sides-to-every-story-bc0780ed7d9c)

Also you can document your database schema using [dbdocs](https://dbdocs.io/).

## TODOS

- Update the config file using [viper](https://github.com/spf13/viper)
- Implemented other use-cases:
  - login user
  - use redis cache for speedup
  - add roles for users
  - ...
