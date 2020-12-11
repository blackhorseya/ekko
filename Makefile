.PHONY: build-image
build-image: gen-wire gen-swagger gen-pb
	docker build -t todo:latest .

.PHONY: run-with-docker
run-with-docker:
	@docker run -it --rm -p 8080:8080 -v $(shell pwd)/configs/app.yaml:/app/configs/app.yaml todo:latest

.PHONY: run-mongo
run-mongo:
	@docker-compose -p todo -f $(shell pwd)/deployments/docker-compose.yml up -d

.PHONY: prune-images
prune-images:
	@docker rmi `docker images --filter=label=app=todo -q`

# todo: 2020-12-11|20:32|doggy|implement me
.PHONY: tag-image

# todo: 2020-12-11|20:31|doggy|implement me
.PHONY: push-image

.PHONY: deploy-with-helm
deploy-with-helm:
	@helm --namespace sean-side-uat-ns upgrade --install todo ./deployments/helm --values ./deployments/helm/values.yaml

.PHONY: gen-pb
gen-pb:
	@protoc --go_out=plugins=grpc:./internal/app/entities --proto_path=./internal/app/protos ./internal/app/protos/*.proto

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
