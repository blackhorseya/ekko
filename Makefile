APP_NAME=todo
VERSION=latest
PROJECT_ID=sean-side
NS=side
DEPLOY_TO=uat
REGISTRY=gcr.io
IMAGE_NAME=$(REGISTRY)/$(PROJECT_ID)/$(APP_NAME)
HELM_REPO_NAME=blackhorseya
CHART_NAME=todo-app

DB_URI="mongodb://todo-app:changeme@localhost:27017/todo-db"

check_defined = $(if $(value $1),,$(error Undefined $1))

.PHONY: clean
clean:
	@rm -rf coverage.txt profile.out ./bin
	@echo Successfuly removed artifacts

.PHONY: test-unit
test-unit:
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: lint
lint:
	@golint ./...

.PHONY: report
report:
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/todo-app'

.PHONY: build-image
build-image:
	$(call check_defined,VERSION)
	@docker build -t $(IMAGE_NAME):$(VERSION) \
	--label "app.name=$(APP_NAME)" \
	--label "app.version=$(VERSION)" \
	--build-arg APP_NAME=$(APP_NAME) \
	--pull --cache-from=$(IMAGE_NAME) \
	-f Dockerfile .

.PHONY: list-images
list-images:
	@docker images --filter=label=app.name=$(APP_NAME)

.PHONY: prune-images
prune-images:
	@docker rmi -f `docker images --filter=label=app.name=$(APP_NAME) -q`

.PHONY: push-image
push-image:
	$(call check_defined,VERSION)
	@docker tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest
	@docker push $(IMAGE_NAME):$(VERSION)
	@docker push $(IMAGE_NAME):latest

.PHONY: up-local-db
up-local-db:
	@docker-compose --file ./deployments/docker-compose.yaml --project-name $(APP_NAME) up -d

.PHONY: down-local-db
down-local-db:
	@docker-compose --file ./deployments/docker-compose.yaml --project-name $(APP_NAME) down -v

.PHONY: deploy
deploy:
	$(call check_defined,VERSION)
	$(call check_defined,DEPLOY_TO)
	@helm --namespace $(NS) \
	upgrade --install $(APP_NAME) $(HELM_REPO_NAME)/$(CHART_NAME) \
	--values ./deployments/configs/$(DEPLOY_TO)/$(APP_NAME).yaml \
	--set image.tag=$(VERSION)

.PHONY: gen
gen: gen-wire gen-swagger gen-pb gen-mocks

.PHONY: gen-pb
gen-pb:
	@protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./pb/*.proto
	@echo Successfully generated proto

.PHONY: gen-wire
gen-wire:
	@wire gen ./...

.PHONY: gen-swagger
gen-swagger:
	@swag init -g cmd/$(APP_NAME)/main.go --parseInternal --parseDependency --parseDepth 1 -o api/docs

.PHONY: gen-mocks # generate mocks code via mockery
gen-mocks:
	@go generate -x ./...

.PHONY: migrate-up
migrate-up:
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migrations up

.PHONY: migrate-down
migrate-down:
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migrations down