APP_NAME = gogolook-assigment
DOCKER_IMAGE = $(APP_NAME):latest

build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

run:
	@echo "Running the application..."
	docker run --rm -p 8888:8888 --name gogolook $(DOCKER_IMAGE)

clean:
	@echo "Cleaning up Docker images..."
	docker rmi $(DOCKER_IMAGE) || true