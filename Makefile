docker-build:
	docker build -t todo:latest .

docker-run:
	docker run -it --rm todo:latest
