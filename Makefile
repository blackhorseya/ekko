build-image:
	docker build -t todo:latest .

run-with-docker:
	@docker run -it --rm -p 8080:8080 -v $(shell pwd)/configs/app.yaml:/app/configs/app.yaml todo:latest

run-mongo:
	@docker-compose -p todo -f $(shell pwd)/deployments/docker-compose.yml up -d

prune-images:
	@docker rmi `docker images --filter=label=app=todo -q`

.PHOYN: deploy-with-helm
deploy-with-helm:
	@helm --namespace sean-side-uat-ns upgrade --install todo ./deployments/helm --values ./deployments/helm/values.yaml

gen-pb:
	@protoc --go_out=plugins=grpc:./internal/app/entities --proto_path=./internal/app/protos ./internal/app/protos/*.proto

gen-wire:
	@wire gen ./...

gen-swagger:
	@swag init -g cmd/app/main.go --parseInternal -o internal/app/docs

test-with-coverage:
	@sh $(shell pwd)/scripts/go.test.sh

lint:
	@golint ./...

install-mod:
	@go mod download

install-tools:
	@go get -v golang.org/x/lint/golint
