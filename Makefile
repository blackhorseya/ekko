docker-build:
	docker build -t todo:latest .

docker-run:
	@docker run -it --rm todo:latest

docker-prune:
	@docker rmi `docker images --filter=label=app=todo -q`
