app_name = todo
app_version = latest
project_id = sean-side-uat
namespace = sean-side-uat-ns

.PHONY: build-image
build-image: gen-wire gen-swagger gen-pb
	@docker build -t $(app_name):$(app_version) \
	--label "app.name=$(app_name)" \
	--label "app.version=$(app_version)" \
	.

.PHONY: list-images
list-images:
	@docker images --filter=label=app.name=$(app_name)

.PHONY: run-with-docker
run-with-docker:
	@docker run -it --rm -p 8080:8080 \
	-v $(shell pwd)/configs/app.yaml:/app/configs/app.yaml \
	$(app_name):$(app_version)

.PHONY: run-mongo
run-mongo:
	@helm --namespace $(namespace) upgrade --install $(app_name)-db bitnami/mongodb \
	--values ./deployments/mongo/local.yaml

.PHONY: prune-images
prune-images:
	@docker rmi -f `docker images --filter=label=app.name=$(app_name) -q`

.PHONY: tag-image
tag-image: build-image
	@docker tag $(app_name):$(app_version) gcr.io/$(project_id)/$(app_name):$(app_version)

.PHONY: push-image
push-image: tag-image
	@docker push gcr.io/$(project_id)/$(app_name):$(app_version)

.PHONY: deploy-with-helm
deploy-with-helm:
	@helm --namespace $(namespace) \
	upgrade --install $(app_name) ./deployments/helm \
	--values ./deployments/helm/values.yaml

.PHONY: gen-pb
gen-pb:
	@protoc --go_out=plugins=grpc:./internal/app/entities \
	--proto_path=./internal/app/protos \
	./internal/app/protos/*.proto

.PHONY: gen-wire
gen-wire:
	@wire gen ./...

.PHONY: gen-swagger
gen-swagger:
	@swag init -g cmd/app/main.go --parseInternal -o internal/app/docs

.PHONY: test-with-coverage
test-with-coverage:
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: lint
lint:
	@golint ./...

.PHONY: install-mod
install-mod:
	@go mod download

.PHONY: install-tools
install-tools:
	@go get -v golang.org/x/lint/golint
