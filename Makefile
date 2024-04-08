include .env

GO_BIN:=$(shell go env GOPATH)/bin

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: build
build: ## build docker image
	@docker build -f ./docker/go/Dockerfile -t server .

.PHONY: run
run: ## run app
	@docker compose up -d 

.PHONY: down
down: ## down app
	@docker compose down --rmi all --volumes --remove-orphans

.PHONY: logs
logs: ## show api server logs
	@docker compose logs -f go-server 

.PHONY: fmt
fmt: ## format go code
	@go fmt ./...

.PHONY: migrate
migrate: ## run migrations
	@migrate -database 'postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_DATABASE)?sslmode=disable' -path migration/ up 1

.PHONY: generate-model
generate-model:
	@$(GO_BIN)/sqlboiler psql

psql: ## login postgresql
	@psql --host localhost --port $(DB_PORT) --username $(DB_USER) --dbname ${DB_DATABASE} --password 
