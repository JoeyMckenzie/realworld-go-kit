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

.PHONY: dev
dev: ## Run the API server
	@go run ./conduit-bin -env development -port 8080 -seed true && ng serve

.PHONY: build
build:  ## Build the API binary
	go build -a -o conduit ./conduit-bin && ng build --omit=dev

.PHONY: clean
clean: ## Remove the application binary
	@rm -f ./conduit && rm -rf ./dist

# TODO: golangci-lint doesn't seem to handle workspace mode properly just yet https://github.com/golangci/golangci-lint/issues/2654
.PHONY: lint
lint: ## Lint all go code
	@golangci-lint run --exclude strings ./source/conduit-api/...
	@golangci-lint run --exclude strings ./source/conduit-bin/...
	@golangci-lint run --exclude strings ./source/conduit-core/...
	@golangci-lint run --exclude strings ./source/conduit-domain/...
	@golangci-lint run --exclude strings ./source/conduit-shared/...

.PHONY: format
format: ## Format all code
	@cd ./source/conduit-api && go fmt ./...
	@cd ./source/conduit-bin && go fmt ./...
	@cd ./source/conduit-core && go fmt ./...
	@cd ./source/conduit-domain && go fmt ./...
	@cd ./source/conduit-shared && go fmt ./...

.PHONY: sync
sync: ## Sync go imports
	@go work sync

.PHONY: test
test: ## Run all tests in the project
	@go test ./source/conduit-api/...
	@go test ./source/conduit-bin/...
	@go test ./source/conduit-core/...
	@go test ./source/conduit-domain/...
	@go test ./source/conduit-shared/...

.PHONY: test-integration
test-integration: ## Runs all integration tests via Postman
	@./integrations/run-postman-tests.sh

.PHONY: start-server
start-server: ## Start the API container
	@docker compose -f ./deploy/docker-compose.server.yml up --build --remove-orphans

.PHONY: start-db
start-db: ## Start the database container
	@docker compose -f ./deploy/docker-compose.postgres.yml up --build --remove-orphans

.PHONY: start-metrics
start-metrics: ## Start the Prometheus metrics container
	@docker compose -f ./deploy/docker-compose.metrics.yml up --build --remove-orphans

.PHONY: start-web
start-web: ## Start the angular web container
	@docker compose -f ./deploy/docker-compose.web.yml up --build --remove-orphans

.PHONY: start-conduit
start-conduit: ## Start all containers required for to run the full application
	@docker compose \
		-f ./deploy/docker-compose.postgres.yml \
 		-f ./deploy/docker-compose.server.yml \
 		-f ./deploy/docker-compose.web.yml \
 		-f ./deploy/docker-compose.metrics.yml up --build

.PHONY: ent-init
ent-init: ## Runs the create entity ent command
	@go run entgo.io/ent/cmd/ent init

.PHONY: ent-generate
ent-generate: ## Generates the ent entity code
	@go generate ./conduit-ent-gen/ent

.PHONY: ent-regenerate
ent-regenerate: ## Regenerates the ent entity code
	@make ent-clean
	@go generate ./conduit-ent-gen/ent

.PHONY: ent-clean
ent-clean: ## Cleans the ent codegen while maintaining existing models
	@./src/conduit-ent-gen/ent-clean.sh
