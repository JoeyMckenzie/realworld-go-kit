![realworld_logo](/realworld-dual-mode.png)

# RealWorld Go Kit

A RealWorld implementation written with Go and go-kit!

This project utilizes several popular Go-based libraries and technologies, primarily including:

- [Go kit](https://gokit.io/) as a general application framework
- [Chi](https://github.com/go-chi/chi) for routing
- [sqlx](https://github.com/jmoiron/sqlx) for querying
- [Fly](https://fly.io/) for hosting
- [Docker](https://www.docker.com/) for containerization
- [PlanetScale](https://planetscale.com/) for database hosting
- [Taskfile](https://taskfile.dev/) for task orchestration
- [golangci-lint](https://golangci-lint.run/) for linting
- [pre-commit](https://pre-commit.com/) for code quality git hooks

## Getting Started

To get started, verify that Docker and Go 1.20 are installed. Next, install the Taskfile CLI and golangci-lint:

```bash
go install github.com/go-task/task/v3/cmd/task@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
```

Next, set your `.env` file:

```bash
mv .env.example .env
```

PlanetScale is based on MySQL, but any old MySQL database will. Run the SQL in the [schema](./schema.sql) file to create
all the necessary tables. If you're using PlanetScale, simply setup an account and add your connection strings to the [.env](./.env) file.

Once the tool are installed, run the application as a task:

```bash
task run
```

That's it! Checkout the [Taskfile](./Taskfile.yml) to see the various tasks available for running. To run integration tests
against the running server, open a new terminal and run:

```bash
task integration:local
```

All tests should pass and there is CI setup to run integration tests against the deployed instance on Fly.
