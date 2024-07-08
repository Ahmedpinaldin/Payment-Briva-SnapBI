include .env
export

export IMAGE_VERSION	  ?=1.18

build:
	docker image rm $(IMAGE_NAME):$(IMAGE_VERSION) --force
	docker build -t $(IMAGE_NAME):$(IMAGE_VERSION) \
	--build-arg PORT=$(PORT) \
	-f ./Dockerfile .

builddev:
	docker image rm $(IMAGE_NAME):$(IMAGE_VERSION) --force
	docker build -t $(IMAGE_NAME):$(IMAGE_VERSION) --build-arg PORT=$(PORT) -f ./Dockerfile1 .

local:
	rm -rf go.sum
	docker run --rm -v $(PWD):/app/ -w /app --name $(IMAGE_NAME) $(IMAGE_NAME):$(IMAGE_VERSION) go version
	docker run --rm -v $(PWD):/app/ -w /app --name $(IMAGE_NAME) $(IMAGE_NAME):$(IMAGE_VERSION) go mod tidy

run:
	docker run -d --name $(IMAGE_NAME) --expose=$(PORT) -p $(PORT):$(PORT) $(IMAGE_NAME):$(IMAGE_VERSION)

dev:
	docker run -it --rm -v $(PWD):/app/ -w /app --name $(IMAGE_NAME) -p $(PORT):$(PORT) --expose=$(PORT) $(IMAGE_NAME):$(IMAGE_VERSION)
