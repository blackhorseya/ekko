APP_NAME=todo-app
VERSION=latest
PROJECT_ID=sean-side
NS=side
DEPLOY_TO=uat
REGISTRY=gcr.io
IMAGE_NAME=$(REGISTRY)/$(PROJECT_ID)/$(APP_NAME)
HELM_REPO_NAME=blackhorseya
CHART_NAME=todo-app

DB_URI='mysql://root:changeme@tcp(localhost:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local'

.PHONY: check-%
check-%: ## check environment variable is exists
	@if [ -z '${${*}}' ]; then echo 'Environment variable $* not set' && exit 1; fi

.PHONY: help
help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean:  ## remove artifacts
	@rm -rf coverage.txt profile.out ./bin
	@echo Successfuly removed artifacts

.PHONY: test-unit
test-unit: ## execute unit test
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: lint
lint: ## execute golint
	@golint ./...

.PHONY: report
report: ## execute goreportcard
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/todo-app'

.PHONY: build-image
build-image: check-VERSION ## build image
	@docker build -t $(IMAGE_NAME):$(VERSION) \
	--label "app.name=$(APP_NAME)" \
	--label "app.version=$(VERSION)" \
	--build-arg APP_NAME=$(APP_NAME) \
	--pull --cache-from=$(IMAGE_NAME) \
	-f Dockerfile .

.PHONY: list-images
list-images: ## list all images
	@docker images --filter=label=app.name=$(APP_NAME)

.PHONY: prune-images
prune-images: ## prune images
	@docker rmi -f `docker images --filter=label=app.name=$(APP_NAME) -q`

.PHONY: push-image
push-image: check-VERSION ## publish image
	@docker tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest
	@docker push $(IMAGE_NAME):$(VERSION)
	@docker push $(IMAGE_NAME):latest

.PHONY: up-local-db
up-local-db: ## run docker-compose up
	@docker-compose --file ./deployments/docker-compose.yaml --project-name $(APP_NAME) up -d

.PHONY: down-local-db
down-local-db: ## run docker-compose down
	@docker-compose --file ./deployments/docker-compose.yaml --project-name $(APP_NAME) down -v

.PHONY: deploy
deploy: check-VERSION check-DEPLOY_TO ## deploy application
	@helm --namespace $(NS) \
	upgrade --install $(DEPLOY_TO)-$(APP_NAME) $(HELM_REPO_NAME)/$(CHART_NAME) \
	--values ./deployments/configs/$(DEPLOY_TO)/values.yaml \
	--set image.tag=$(VERSION)

.PHONY: gen
gen: gen-wire gen-pb gen-mocks gen-swagger ## generate code

.PHONY: gen-pb
gen-pb: ## generate protobuf
	@protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./pb/*.proto
	@echo Successfully generated proto

.PHONY: gen-wire
gen-wire: ## generate wire
	@wire gen ./...

.PHONY: gen-swagger
gen-swagger: ## generate swagger spec
	@swag init -g cmd/$(APP_NAME)/main.go --parseInternal --parseDependency --parseDepth 1 -o api/docs

.PHONY: gen-mocks # generate mocks code via mockery
gen-mocks: ## generate mocks
	@go generate -tags=wireinject -x ./...

.PHONY: migrate-up
migrate-up: ## run migration up
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migrations up

.PHONY: migrate-down
migrate-down: ## run migration down
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migrations down

.PHONY: update-package
update-package: ## update package and commit
	@go get -u ./...
	@go mod tidy
	@git add go.mod go.sum
	@git commit -m "build: update package"
