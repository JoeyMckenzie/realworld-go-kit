MODULE = $(shell go list -m)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo "1.0.0")

ENV_FILE ?= .env
DSN ?= $(shell sed -n 's/^CONNECTION_STRING=\(.*\)/\1/p' $(ENV_FILE))

.PHONY: default
default: help

# generate help info from comments: thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version: ## Displays the version of the API server
	@echo $(VERSION)

.PHONY: start
start: ## Run the API server
	@go run ./cmd/api -env development -port 8080 -seed true

.PHONY: build
build:  ## Build the API binary
	go build -a -o conduit ./cmd/api

.PHONY: clean
clean: ## Remove the application binary
	@rm -f conduit

.PHONY: lint
lint: ## Lint all go code
	@golangci-lint run --exclude strings

.PHONY: format
format: ## Format all code
	@go fmt ./...

.PHONY: tidy
tidy: ## Tidy go imports
	@go mod tidy

.PHONY: test
test: ## Run all tests in the project
	@go test ./...

.PHONY: test-integration
test-integration: ## Runs all integration tests via Postman
	@./run-postman-tests

.PHONY: start-api
start-api: ## Start the API container
	@docker compose -f ./docker-compose.api.yml up --build

.PHONY: start-db
start-db: ## Start the database container
	@docker compose -f ./docker-compose.postgres.yml up --build --remove-orphans

.PHONY: start-metrics
start-metrics: ## Start the Prometheus metrics container
	@docker compose -f ./docker-compose.metrics.yml up --build

.PHONY: start-conduit
start-conduit: ## Start all containers required for to run the full application
	@docker compose -f ./docker-compose.postgres.yml -f ./docker-compose.api.yml -f ./docker-compose.metrics.yml up --build

.PHONY: install-deps
install-deps: ## Installs all application package dependencies
	go get github.com/go-chi/chi/v5
	go get github.com/go-chi/cors
	go get github.com/joho/godotenv
	go get github.com/lib/pq
	go get github.com/go-kit/log
	go get github.com/go-kit/kit/endpoint
	go get github.com/go-kit/kit/transport/http
	go get github.com/go-kit/kit/metrics/prometheus
	go get github.com/prometheus/client_golang/prometheus
	go get github.com/go-playground/validator/v10
	go get github.com/golang-jwt/jwt
	go get golang.org/x/crypto
	go get github.com/gosimple/slug
	go get github.com/mattn/go-sqlite3

.PHONY: ent-init
ent-init: ## Runs the create entity ent command
	@go run entgo.io/ent/cmd/ent init

.PHONY: ent-generate
ent-generate: ## Generates the ent entity code
	@go generate ./conduit-ent-gen/ent

.PHONY: ent-regenerate
ent-regenerate: ## Generates the ent entity code
	@make ent-clean
	@go generate ./conduit-ent-gen/ent

.PHONY: ent-clean
ent-clean: ## Cleans the ent codegen while maintaining existing models
	@cd ./conduit-ent-gen
	@mkdir ./conduit-ent-gen/tmp
	@cp -r ./conduit-ent-gen/ent/schema ./conduit-ent-gen/tmp
	@cp ./conduit-ent-gen/ent/generate.go ./conduit-ent-gen/tmp
	@rm -rf ./conduit-ent-gen/ent
	@mkdir ./conduit-ent-gen/ent
	@cp -r ./conduit-ent-gen/tmp/schema ./conduit-ent-gen/ent
	@cp ./conduit-ent-gen/tmp/generate.go ./conduit-ent-gen/ent
	@rm -rf ./conduit-ent-gen/tmp
