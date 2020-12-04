build-image:
	docker build -t todo:latest .

run-with-docker:
	@docker run -it --rm todo:latest

prune-images:
	@docker rmi `docker images --filter=label=app=todo -q`

gen-pb:
	@protoc --go_out=plugins=grpc:./internal/entities --proto_path=./internal/protos ./internal/protos/*.proto

test-with-unit:
	@go test -v ./...

test-with-coverage:
	@go test -race -coverprofile=coverage.txt -covermode=atomic ./...

install-mod:
	@go mod download

install-tools:
	@go get -v golang.org/x/lint/golint
