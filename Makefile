build-image:
	docker build -t todo:latest .

run-with-docker:
	@docker run -it --rm -p 8080:8080 -v $(shell pwd)/configs/app.yaml:/app/configs/app.yaml todo:latest

prune-images:
	@docker rmi `docker images --filter=label=app=todo -q`

gen-pb:
	@protoc --go_out=plugins=grpc:./internal/app/entities --proto_path=./internal/app/protos ./internal/app/protos/*.proto

gen-wire:
	@wire gen ./cmd/...

test-with-unit:
	@go test -race -v ./...

test-with-coverage:
	@go test -race -cover -coverpkg=./... -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	@golint ./...

install-mod:
	@go mod download

install-tools:
	@go get -v golang.org/x/lint/golint
