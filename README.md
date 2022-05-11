![realworld_logo](/realworld-dual-mode.png)

This project is an implementation of the RealWorld project using Go utilizing [go-kit](https://gokit.io/) as an
application framework.
The project heavily utilizes Docker for containerization of dependencies as well as containerizing the API itself.
To get started, make sure you have [Docker](https://www.docker.com/) and [make](https://www.gnu.org/software/make/)
installed on your machine, then clone/fork the repository. Once you have the project
local to your machine, create an `.env` file local to workspace with valid values (you can use the defaults as well).

```bash
mv .env.example .env
make start-conduit
```

Once the application containers have started, verify all integration tests pass:

```
make test-integration
```

The above target will run the included Postman suite of tests designed by the authors of the RealWorld project.
Once the tests have completed, verify all unit tests are passing as well:

```
make test
```

Again, the target above will run all included unit tests found in the project.

## Application architecture

As mentioned before, the project utilizes go-kit as the application framework. In general, an application feature is
composed of several parts:

- Domain - an aggregation of plain ol' Go structs describing the shapes of requests and responses used to transport
  pieces of data between layers
- API - the outermost layer comprised of go-kit transports and endpoints responsible for receiving and responding to
  HTTP requests
- Core - a feature core contains services and utilities housing all business logic

The project uses a handful of libraries, namely:

- [go-kit](https://gokit.io/) as a framework for HTTP transports, middlewares, logging, etc.
- [ent](https://entgo.io/) as an ORM and persistence API
- [chi](https://github.com/go-chi/chi) for routing
- [godotenv](https://github.com/joho/godotenv) for using `.env` files
- [pq](https://github.com/lib/pq) as a Postgres database driver
- [prometheus](https://github.com/prometheus/client_golang/prometheus) for service metrics
- [validator](https://github.com/go-playground/validator/v10) for request struct validation
- [jwt](https://github.com/golang-jwt/jwt) for signing, creating, and validating JWTs
- [crypto](https://golang.org/x/crypto) for password hashing
- [slug](https://github.com/gosimple/slug) for slugifying article titles
- [sqlite3](https://github.com/mattn/go-sqlite3) as an in-memory database for unit and integration testing
- [stretchr](https://github.com/stretchr/testify) for test assertions and mocking

#### Why are you using a `Makefile`, aren't there better options out there?

Certainly! The project started off using [taskfile](https://taskfile.dev/#/), but I made the choice
to swap to make purely out of preference and really nothing else. For starters, I like make as a standard project
toolchain
of sorts so that I don't have to memorize a languages native build tool. I use [GoLand](https://www.jetbrains.com/go/)
as my IDE of choice
and the folks at JetBrains have wonderful integrations with Makefiles that make developing `make`-based projects a
breeze.

#### Why are you using go-kit?

Among many things, go-kit offers convention as a framework for building Go applications. Go-kit includes all the things
us
developers love when spinning up new services, namely built-in structured logging, metrics that integrate with
prometheus out of the box,
and a clear and concise API for wiring up endpoints to services with middleware to boot. While go-kit can be a bit
boilerplate-y, I quite
enjoy the convention that go-kit and plan on using it for future applications.

#### Why are you using ent, aren't ORMs bad?

That depends. There are certainly sound arguments on both sides of the fence, and the choice of tooling for persistence
should
be taken into careful consideration before bringing in a core project dependency. I initially opted
for [sqlx and atlas](https://github.com/JoeyMckenzie/realworld-go-kit/tree/archive/sqlx),
but wanted something a bit more developer friendly and easy to use so that I could focus on writing the core business
logic
rather than writing SQL. Writing SQL is not hard - writing good SQL is hard.

Secondarily, as a .NET developer, ent offers a similar API
to [Entity Framework](https://docs.microsoft.com/en-us/ef/core/) that us
.NET developers have been using for years
and (mostly) know and love. Being also a big fan of [Dapper](https://github.com/DapperLib/Dapper), sqlx was the obvious
choice starting out, but I opted for ent as it seemed fun and wanted to give it a try.

#### Why is there `package.json` file?

I use [husky](https://github.com/typicode/husky) for pre-commit hooks
and [lint-staged](https://www.npmjs.com/package/lint-staged)
to format staged files to keep committed code well formatted. While there are a few other options for including
pre-commit hooks
into a Go project, and certainly those that are more appropriate for Go projects, I wanted to leave open the opportunity
of bringing on a JS-based frontend sometime in the future to have the true RealWorld fullstack experience. The
pre-commit hooks will format, lint, and test all code so that each commit ensure that tests are passing and code does
not contain
any obvious errors.

## Using Docker

The project utilizes Docker containers for Postgres and prometheus metrics. For example, when starting the
application with `make start-conduit`, navigating to `localhost:9090` will bring you to the prometheus metrics page.
From there,
running integration tests with `make test-integration` to simulate traffic to the API will allow you the various metrics
that are recorded in the service layer: request count, request latency, and histograms of service request intervals.

To start the API outside the Docker context, ensure that Postgres is running before booting up:

```bash
make start-db
```

Once the Postgres container has started, go ahead and spin up the API for active development:

```bash
make start
```

If you're starting the application for the first time, it will attempt to seed a bit of data that is also used for
testing.

## Todo

There's quite a bit of code cleanup to do throughout the project, and any and all issues/PRs are welcome. Good first issues
would be helping with code cleanup and adding more unit test coverage. Eventually, I would like to include a sweet of integration tests using `httptest` as well.
