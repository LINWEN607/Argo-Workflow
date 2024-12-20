# Makefile

# Variables
IMAGE_NAME ?= go-gin:v1.0  # 指定镜像名称，默认为 your-image-name
COMPOSE_FILE ?= docker-compose.yaml  # 指定 docker-compose 文件，默认为 docker-compose.yaml

gotest:
	@go test ./... -v -cover
# Build the Docker image
build:
	@docker build -t $(IMAGE_NAME) .

# Run the Docker Compose services in detached mode
run:
	@docker-compose -f $(COMPOSE_FILE) up -d

# Phony targets
.PHONY: build run
