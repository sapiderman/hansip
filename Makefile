GOPATH=$(shell go env GOPATH)
IMAGE_REGISTRY=dockerhub
IMAGE_NAMESPACE ?= hansip
IMAGE_NAME ?= $(shell basename `pwd`)
CURRENT_PATH=$(shell pwd)
COMMIT_ID ?= $(shell git rev-parse --short HEAD)
GO111MODULE=on

.PHONY: all test clean build docker

build-static:
	-${GOPATH}/bin/go-resource -base "$(CURRENT_PATH)/api/swagger-ui" -path "/docs" -filter "/**/*" -go "$(CURRENT_PATH)/api/StaticApi.go" -package api
	go fmt ./...

build: build-static
	export GO111MODULE=on; \
	GO_ENABLED=0 go build -a -o $(IMAGE_NAME).app cmd/main/Main.go
#   Use bellow if you're running on linux.
#	GO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o $(IMAGE_NAME).app cmd/main/Main.go

lint: build-static
	golint -set_exit_status ./internal/... ./pkg/...

test: build-static
	export GO111MODULE on; \
	go test ./... -cover -vet -all -v -short

run: build
	export AAA_SERVER_HOST=0.0.0.0; \
	export AAA_SERVER_PORT=8088; \
	export AAA_SETUP_ADMIN_ENABLE=true; \
	export AAA_SERVER_LOG_LEVEL=TRACE; \
	export AAA_DB_TYPE=MYSQL; \
	./$(IMAGE_NAME).app
	rm -f $(IMAGE_NAME).app

docker:
	docker build -t $(IMAGE_NAMESPACE)/$(IMAGE_NAME):latest -f ./.docker/Dockerfile .

docker-build-commit: build
	docker build -t $(IMAGE_NAMESPACE)/$(IMAGE_NAME):$(COMMIT_ID) -f ./.docker/Dockerfile .

docker-build: build
	docker build -t $(IMAGE_NAMESPACE)/$(IMAGE_NAME):$(COMMIT_ID) -f ./.docker/Dockerfile .
	docker tag $(IMAGE_NAMESPACE)/$(IMAGE_NAME):$(COMMIT_ID) $(IMAGE_NAMESPACE)/$(IMAGE_NAME):latest

docker-push:
	docker push $(IMAGE_NAMESPACE)/$(IMAGE_NAME):$(COMMIT_ID)

docker-stop:
	-docker stop $(IMAGE_NAME)

docker-rm: docker-stop
	-docker rm $(IMAGE_NAME)

docker-run: docker-rm docker
	docker run --name $(IMAGE_NAME) -p 3000:3000 --detach $(IMAGE_NAMESPACE)/$(IMAGE_NAME):latest
