![realworld_logo](/realworld-dual-mode.png)

This project is an implementation of the RealWorld project using Go utilizing [go-kit](https://gokit.io/) as an application framework.
The project heavily utilizes Docker for containerization of dependencies as well as containerizing the API itself.
To get started, make sure you have [Docker](https://www.docker.com/) and [make](https://www.gnu.org/software/make/) installed on your machine, then clone/fork the repository. Once you have the project
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

As mentioned before, the project utilizes go-kit as the application framework. In general, an application feature is composed of several parts:
- Domain - an aggregation of plain ol' Go structs describing the shapes of requests and responses used to transport pieces of data between layers 
- API - the outermost layer comprised of go-kit transports and endpoints responsible for receiving and responding to HTTP requests
- Core - a feature core contains services and utilities housing all business logic