# env for project
PROJECT_NAME := ekko
VERSION := $(shell git describe --tags --always)

## common
.PHONY: help
help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version: ## show version
	@echo $(VERSION)

.PHONY: report
report: ## execute goreportcard
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/ekko'

.PHONY: clean
clean:  ## remove artifacts
	@rm -rf cover.out result.json ./bin ./deployments/charts/*.tgz
	@echo Successfuly removed artifacts

## go
.PHONY: lint
lint: ## run golangci-lint
	@golangci-lint run ./...

.PHONY: gazelle
gazelle: ## run gazelle with bazel
	@bazel run //:gazelle

.PHONY: build
build: ## build go binary
	@bazel build //...

.PHONY: test
test: ## test go binary
	@bazel test //...

.PHONY: coverage
coverage: ## generate coverage report
	@go test -json -coverprofile=cover.out ./... >result.json

.PHONY: gen-swagger
gen-swagger: ## generate swagger
	@@swag init -q -d ./adapter,./entity,./pkg -o ./adapter/api

## docker
.PHONY: docker-push
docker-push: ## push docker image
	@bazel run //adapter:push --platforms=@rules_go//go/toolchain:linux_amd64 -- --tag=$(VERSION)
