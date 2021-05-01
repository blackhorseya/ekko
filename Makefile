app_name = todo
version = latest
project_id = sean-side
ns = side
deploy_to = uat

.PHONY: clean
clean:
	@rm -rf coverage.txt profile.out bin

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
	@docker build -t $(app_name):$(version) \
	--label "app.name=$(app_name)" \
	--label "app.version=$(version)" \
	--pull .

.PHONY: list-images
list-images:
	@docker images --filter=label=app.name=$(app_name)

.PHONY: prune-images
prune-images:
	@docker rmi -f `docker images --filter=label=app.name=$(app_name) -q`

.PHONY: tag-image
tag-image:
	@docker tag $(app_name):$(version) gcr.io/$(project_id)/$(app_name):$(version)

.PHONY: push-image
push-image:
	@docker push gcr.io/$(project_id)/$(app_name):$(version)

.PHONY: install-db
install-db:
	@helm --namespace $(ns) upgrade --install $(app_name)-db bitnami/mongodb \
	--values ./deployments/configs/$(deploy_to)/mongo.yaml

.PHONY: deploy
deploy:
	@helm --namespace $(ns) \
	upgrade --install $(app_name) ./deployments/helm \
	--values ./deployments/configs/$(deploy_to)/todo.yaml \
	--set image.tag=$(version)

.PHONY: gen
gen: gen-pb gen-swagger gen-wire

.PHONY: gen-pb
gen-pb:
	@protoc --go_out=plugins=grpc:./internal/pkg/entity ./internal/pkg/entity/**/*.proto

.PHONY: gen-wire
gen-wire:
	@wire gen ./...

.PHONY: gen-swagger
gen-swagger:
	@swag init -g cmd/app/main.go --parseInternal -o api/docs
