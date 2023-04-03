# env for project
APP_NAME := ekko
VERSION := $(shell git describe --tags --always)
SVC_NAME := task
SVC_ADAPTER := restful
MAIN_PKG := ./cmd/$(SVC_ADAPTER)/$(SVC_NAME)

# env for gcp
PROJECT_ID := $(shell gcloud config get-value project)
REGISTRY := gcr.io
IMAGE_NAME := $(REGISTRY)/$(PROJECT_ID)/$(APP_NAME)-$(SVC_NAME)-$(SVC_ADAPTER)

# env for helm
HELM_REPO_NAME := sean-side

# env for k8s
DEPLOY_TO := prod
NS := $(APP_NAME)
RELEASE_NAME := $(DEPLOY_TO)-$(APP_NAME)-$(SVC_NAME)-$(SVC_ADAPTER)

## common
.PHONY: check-%
check-%: ## check environment variable is exists
	@if [ -z '${${*}}' ]; then echo 'Environment variable $* not set' && exit 1; fi

.PHONY: help
help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: report
report: ## execute goreportcard
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/ekko'

.PHONY: clean
clean:  ## remove artifacts
	@rm -rf coverage.txt profile.out ./bin ./deployments/charts/*.tgz
	@echo Successfuly removed artifacts

## go
.PHONY: test-unit
test-unit: ## execute unit test
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: test-e2e
test-e2e: ## execute e2e test
	@cd ./test/e2e && npx playwright test ./tests

.PHONY: lint
lint: ## execute golint
	@golint ./...

## docker
.PHONY: build-image
build-image: ## build image
	@echo "Building image $(IMAGE_NAME):$(VERSION)"
	@docker build -t $(IMAGE_NAME):$(VERSION) \
	--label "app.name=$(APP_NAME)" \
	--label "app.version=$(VERSION)" \
	--build-arg MAIN_PKG=$(MAIN_PKG) \
	--pull --cache-from $(IMAGE_NAME):latest \
	--platform linux/amd64 \
	--pull -f Dockerfile .

.PHONY: list-image
list-image: ## list all images
	@docker images --filter=label=app.name=$(APP_NAME)

.PHONY: prune-image
prune-image: ## prune images
	@docker rmi -f `docker images --filter=label=app.name=$(APP_NAME) -q`

.PHONY: push-image
push-image: ## publish image
	@echo "Pushing image $(IMAGE_NAME):$(VERSION)"
	@docker tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest
	@docker push $(IMAGE_NAME):$(VERSION)
	@docker push $(IMAGE_NAME):latest

## generate
.PHONY: gen
gen: gen-wire gen-pb gen-mocks gen-swagger ## generate code

.PHONY: gen-pb
gen-pb: ## generate protobuf messages and services
	@go get -u google.golang.org/protobuf/proto
	@go get -u google.golang.org/protobuf/cmd/protoc-gen-go

	## Starting generate pb
	@protoc --proto_path=. \
			--go_out=. --go_opt=module=github.com/blackhorseya/ekko \
			--go-grpc_out=. --go-grpc_opt=module=github.com/blackhorseya/ekko,require_unimplemented_servers=false \
			./pb/domain/*/**.proto
	@echo Successfully generated proto

	## Starting inject tags
	@protoc-go-inject-tag -input="./pkg/entity/domain/*/model/*.pb.go"
	@echo Successfully injected tags

.PHONY: gen-wire
gen-wire: ## generate wire
	@wire gen ./...

.PHONY: gen-swagger
gen-swagger: ## generate swagger spec
	@swag init -q --dir ./cmd/restful/task,./ -o ./api/docs
	## Generated swagger spec

.PHONY: gen-mocks
gen-mocks: ## generate mocks
	@go generate -tags=wireinject -x ./...

.PHONY: gen-build
gen-build: ## run gazelle with bazel
	@bazel run //:gazelle

DB_URI='mysql://root:changeme@tcp(localhost:3306)/ekko?charset=utf8mb4&parseTime=True&loc=Local'
N=1

## database
.PHONY: migrate-up
migrate-up: ## run migration up
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migrations up

.PHONY: migrate-down
migrate-down: ## run migration down
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migrations down $(N)

## dependency
.PHONY: update-package
update-package: ## update package and commit
	@go get -u ./...
	@go mod tidy

	@bazel run //:gazelle-update-repos

	@git add go.mod go.sum deps.bzl
	@git commit -m "build: update package"

## helm
.PHONY: lint-helm
lint-helm: ## lint helm chart
	@helm lint deployments/charts/*

.PHONY: add-helm-repo
add-helm-repo: ## add helm repo
	@helm repo add --no-update $(HELM_REPO_NAME) gs://sean-helm-charts/charts
	@helm repo update $(HELM_REPO_NAME)

.PHONY: package-helm
package-helm: ## package helm chart
	@helm package ./deployments/charts/$(APP_NAME) --destination ./deployments/charts

.PHONY: push-helm
push-helm: ## push helm chart to gcs
	@helm gcs push --force ./deployments/charts/$(APP_NAME)-*.tgz $(HELM_REPO_NAME)
	@helm repo update $(HELM_REPO_NAME)

.PHONY: upgrade-helm
upgrade-helm: ## upgrade helm chart
	@echo "Upgrading $(RELEASE_NAME) to $(VERSION)"
	@helm upgrade $(RELEASE_NAME) $(HELM_REPO_NAME)/$(APP_NAME) \
	--install --namespace $(NS) --create-namespace \
	--history-max 3 \
	--values ./deployments/configs/$(DEPLOY_TO)/values.yaml \
	--set image.tag=$(VERSION)
