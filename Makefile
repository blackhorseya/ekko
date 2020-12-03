build-image:
	docker build -t todo:latest .

run-with-docker:
	@docker run -it --rm todo:latest

prune-images:
	@docker rmi `docker images --filter=label=app=todo -q`
