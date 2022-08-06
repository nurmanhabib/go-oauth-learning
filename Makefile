PROJECT_NAME := "go-oauth2"
PKG := "gitlab.com/nurmanhabib/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
BINARY_NAME := "go-oauth2-service"

.PHONY: dep lint critic cyclo unit-test integration-test race-test report-test build clean secret hooks migrate seed grpc docker-build docker-compose-build

## Docker:
docker-build: ## Use the dockerfile to build the container
	docker build -f deploy/Dockerfile --build-arg userID=$(userID) --build-arg secretID=$(secretID) --tag ${BINARY_NAME} .

docker-compose-build: ## Start all or c=<name> containers in foreground
	docker-compose -f deploy/docker-compose.yml build --no-cache --build-arg userID=$(userID) --build-arg secretID=$(secretID) --build-arg version=$(BUILD_VERSION)
	docker-compose -f deploy/docker-compose.yml $(c)

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

critic: ## Critic the files
	@gocritic check ${PKG_LIST}

cyclo: ## Cyclo detection over 15 degree complexity
	@gocyclo -over 15 .

dep: ## Get the dependencies
	@go get -v ./...

integration-test: dep ## Run integration tests
	@go test -v ${PKG_LIST} -p 1 -cover -coverprofile=coverage.out

unit-test: dep ## Run unit tests
	@go test -v -short ${PKG_LIST} -cover -coverprofile=coverage.out

race-test: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

report-test:
	@go tool cover -html=coverage.out

build: dep ## Build the binary file
	@go build -i -v $(PKG)

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

secret: ## Create application secret
	@go run main.go create:secret

hooks: ## Init git hooks
	@cp .githooks/pre-commit .git/hooks
	@chmod +x .git/hooks/pre-commit
	@git config core.hooksPath .git/hooks

migrate: ## Init db migration
	@go run main.go db:migrate

seed: ## Init db seed
	@go run main.go db:init

grpc: ## Start grpc server
	@go run main.go grpc:start

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
