## Content
- [Quick start](#quick-start)
- [Project structure](#project-structure)

## Quick start
Local development:
```sh
# Postgres, RabbitMQ
$ make compose-up
# Run app with migrations
$ make run
```

Integration tests (can be run in CI):
```sh
# DB, app + migrations, integration tests
$ make compose-up-integration-test
```

## Project structure
### `cmd/app/main.go`
Configuration and logger initialization. Then the main function "continues" in
`internal/app/app.go`.

### `config`
Configuration. First, `config.yml` is read, then environment variables overwrite the yaml config if they match.
The config structure is in the `config.go`.

### `docs`
Swagger documentation. Auto-generated by [swag](https://github.com/swaggo/swag) library.

### `integration-test`
Integration tests.
They are launched as a separate container, next to the application container.
It is convenient to test the Rest API using [go-hit](https://github.com/Eun/go-hit).

### `internal/app`
There is always one _Run_ function in the `app.go` file, which "continues" the _main_ function.
This is where all the main objects are created.

The `migrate.go` file is used for database auto migrations.
It is included if an argument with the _migrate_ tag is specified.
For example:

```sh
$ go run -tags migrate ./cmd/app
```

### `internal/controller`
Server handler layer (MVC controllers). The template shows 2 servers:
- RPC (RabbitMQ as transport)
- REST http (Gin framework)

Server routers are written in the same style:
- Handlers are grouped by area of application (by a common basis)
- For each group, its own router structure is created, the methods of which process paths
- The structure of the business logic is injected into the router structure, which will be called by the handlers

#### `internal/controller/http`
Simple REST versioning.

In `v1/router.go` and above the handler methods, there are comments for generating swagger documentation using [swag](https://github.com/swaggo/swag).

### `internal/entity`
Entities of business logic (models) can be used in any layer.
There can also be methods, for example, for validation.

### `internal/usecase`
Business logic.
- Methods are grouped by area of application (on a common basis)
- Each group has its own structure
- One file - one structure

Repositories, rpc, and other business logic structures are injected into business logic structures

#### `internal/usecase/repo`
A repository is an abstract storage (database) that business logic works with.

### `pkg/rabbitmq`
RabbitMQ RPC pattern:
- There is no routing inside RabbitMQ
- Exchange fanout is used, to which 1 exclusive queue is bound, this is the most productive config
- Reconnect on the loss of connection

--- 
Created from [Go-clean-template](https://evrone.com/go-clean-template?utm_source=github&utm_campaign=go-clean-template) which is created & supported by [Evrone](https://evrone.com/?utm_source=github&utm_campaign=go-clean-template).