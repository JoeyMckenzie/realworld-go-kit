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
	@go run ./conduit-bin -env development -port 8080 -seed true

.PHONY: build
build:  ## Build the API binary
	go build -a -o conduit ./conduit-bin

.PHONY: clean
clean: ## Remove the application binary
	@rm -f ./conduit

.PHONY: lint
lint: ## Lint all go code
	golangci-lint run --exclude strings ./conduit-*

.PHONY: format
format: ## Format all code
	@go fmt ./conduit-*

.PHONY: sync
sync: ## Sync go imports
	@go work sync

.PHONY: test
test: ## Run all tests in the project
	@go test ./conduit-*

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
