# env for project
PROJECT_NAME := ekko
APP_NAME := ekko
VERSION := $(shell git describe --tags --abbrev=0)
DOMAIN_NAME := issue

# env for k8s
DEPLOY_TO := prod
NS := $(APP_NAME)
RELEASE_NAME := $(DEPLOY_TO)-$(APP_NAME)

## common
.PHONY: check-%
check-%: ## check environment variable is exists
	@if [ -z '${${*}}' ]; then echo 'Environment variable $* not set' && exit 1; fi

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
	@rm -rf coverage.txt profile.out ./bin ./deployments/charts/*.tgz
	@echo Successfuly removed artifacts

## go
.PHONY: test-e2e
test-e2e: ## execute e2e test
	@cd ./test/e2e && npx playwright test ./tests

.PHONY: lint
lint: ## execute golint
	@golint ./...

.PHONY: gazelle-repos
gazelle-repos: ## update gazelle repos
	@bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies -prune

.PHONY: gazelle
gazelle: gazelle-repos ## run gazelle with bazel
	@bazel run //:gazelle

.PHONY: build-go
build-go: ## build go binary
	@bazel build //...

.PHONY: test-go
test-go: ## test go binary
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: push-image
push-image: ## push image to gcr
	@bazel run //adapter/restful:push-image

## generate
.PHONY: gen
gen: gen-pb-go gen-mocks gen-swagger ## generate code

.PHONY: gen-pb-go
gen-pb-go: ## generate go protobuf
	@go get -u google.golang.org/protobuf/proto
	@go get -u google.golang.org/protobuf/cmd/protoc-gen-go

	## Starting generate pb
	@protoc --proto_path=./pb --go_out=paths=source_relative:./entity --go-grpc_out=paths=source_relative,require_unimplemented_servers=false:./pb ./pb/domain/*/*/*.proto
	@echo Successfully generated proto

	## Starting inject tags
	@protoc-go-inject-tag -input="./entity/domain/*/model/*.pb.go"
	@echo Successfully injected tags

.PHONY: gen-swagger
gen-swagger: ## generate swagger spec
	@swag init -q -d ./adapter/restful,./pkg,./entity/domain/task/model -o ./adapter/restful/api/docs
	## Generated swagger spec

.PHONY: gen-mocks
gen-mocks: ## generate mocks
	## Starting generate wire and mockgen
	@go generate -tags="wireinject" ./...
	@echo Successfully generated wire and mockgen

## database
DB_RELEASE_NAME := $(DEPLOY_TO)-$(APP_NAME)-$(DOMAIN_NAME)-db
DB_URI := 'mysql://root:changeme@tcp(localhost:3306)/ekko?charset=utf8mb4&parseTime=True&loc=Local'
N := 1

.PHONY: upgrade-db
upgrade-db: ## upgrade database
	@echo "Deploying database $(DB_RELEASE_NAME) to $(DEPLOY_TO)"
	@helm upgrade $(DB_RELEASE_NAME) bitnami/mariadb \
	--install --namespace $(NS) --create-namespace \
	--history-max 3 \
	--values ./deployments/configs/$(DEPLOY_TO)/db.yaml

.PHONY: migrate-up
migrate-up: ## run migration up
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migrations up

.PHONY: migrate-down
migrate-down: ## run migration down
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migrations down $(N)

## helm
HELM_REPO_NAME := sean-side

.PHONY: lint-helm
lint-helm: ## lint helm chart
	@helm lint deployments/charts/*

.PHONY: install-helm-plugins
install-helm-plugins: ## install helm plugins
	@helm plugin install https://github.com/hayorov/helm-gcs.git

.PHONY: add-helm-repo
add-helm-repo: ## add helm repo
	@helm repo add --no-update $(HELM_REPO_NAME) gs://sean-helm-charts/charts
	@helm repo update $(HELM_REPO_NAME)

.PHONY: package-helm
package-helm: ## package helm chart
	@helm package ./deployments/charts/* --destination ./deployments/charts

.PHONY: push-helm
push-helm: ## push helm chart to gcs
	@for file in $(wildcard ./deployments/charts/*.tgz); do \
		filename=$$(basename $$file); \
		helm gcs push --force $$file $(HELM_REPO_NAME); \
	done
	@helm repo update $(HELM_REPO_NAME)

.PHONY: upgrade-helm
upgrade-helm: ## upgrade helm chart
	@echo "Upgrading $(RELEASE_NAME) to $(VERSION)"
	@helm upgrade $(RELEASE_NAME) $(HELM_REPO_NAME)/$(APP_NAME) \
	--install --namespace $(NS) --create-namespace \
	--history-max 3 \
	--values ./deployments/configs/$(DEPLOY_TO)/$(APP_NAME).yaml
