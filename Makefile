GO111MODULES    :=on
APP             :=summer
REGISTRY?       :=gcr.io/images
COMMIT_SHA      :=$(shell git rev-parse --short HEAD)
CMD_DIR         :=cmd
BIN_DIR         :=bin
REGISTRY        :=sekulicd/summer_repo
ENV             :=dev
TAG             :=latest
GRPC_CLIENT_DIR :=summer-cli
GRPC_SERVER_DIR :=summer-server


#### CONTINUOUS INTEGRATION ###

#ci: clean vet build test cov
ci: clean vet build test cov

#cd
cd: docker-login docker-build docker-tag docker-push

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning..."
	@go clean

.PHONY: vet
## vet: code analysis
vet:
	@echo "Vet..."
	@go vet ./...

.PHONY: build
## build: build the application
build:
	@echo "Building..."
	CGO_ENABLED=0 GOOS=linux go build -o ${BIN_DIR}/${GRPC_CLIENT_DIR} $(CMD_DIR)/${GRPC_CLIENT_DIR}/main.go
	CGO_ENABLED=0 GOOS=linux go build -o ${BIN_DIR}/${GRPC_SERVER_DIR} $(CMD_DIR)/${GRPC_SERVER_DIR}/main.go

.PHONY: test
## test: runs go test with default values
test:
	@echo "Testing..."
	go test -v -count=1 -race ./...

.PHONY: cov
## cov: generates coverage report
cov:
	@echo "Coverage..."
	go test -cover ./...

#### CONTINUOUS DELIVERY/DEPLOYMENT ###

.PHONY: docker-login
## docker-build: builds the summer docker image to registry
docker-login:
	docker login --username=${DOCKER_USERNAME} --password=${DOCKER_PASSWORD}

.PHONY: docker-build
## docker-build: builds the summer docker image to registry
docker-build:
	docker build -t ${APP} .

.PHONY: docker-tag
## docker-tag: run the summer docker container
docker-tag:
	docker tag ${APP} ${REGISTRY}/${TAG}

.PHONY: docker-push
## docker-push: pushes the summer docker image to registry
docker-push: check-environment docker-tag
	docker push ${REGISTRY}:${TAG}


#### LOCAL DEV

.PHONY: build-local
## build: build the application
build-local:
	@echo "Building..."
	go build -o ${BIN_DIR}/${GRPC_CLIENT_DIR} $(CMD_DIR)/${GRPC_CLIENT_DIR}/main.go
	go build -o ${BIN_DIR}/${GRPC_SERVER_DIR} $(CMD_DIR)/${GRPC_SERVER_DIR}/main.go

.PHONY: install
## install-local: install CLI
install:
	@echo "Installing..."
	go install ./${CMD_DIR}/${GRPC_CLIENT_DIR}

.PHONY: run-server
## run: runs go run main.go
run-server:
	go run -race $(CMD_DIR)/${GRPC_SERVER_DIR}/main.go

.PHONY: setup
## setup: setup go modules
setup:
	@go mod init \
		&& go mod tidy \
		&& go mod vendor

.PHONY: download
## setup: download go modules
download:
	@go mod download

.PHONY: grpc-gen
## grpc-gen: generate pb.gp
grpc-gen:
	@protoc --go_out=plugins=grpc:. ./pkg/grpc-schema/summer.proto

# helper rule for deployment
check-environment:
ifndef ENV
	$(error ENV not set, allowed values - `staging` or `production`)
endif

.PHONY: docker-run
## docker-run: run the summer docker container
docker-run: docker-build
	docker run -p8080:3000 -d ${APP}

.PHONY: help
## help: Prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
