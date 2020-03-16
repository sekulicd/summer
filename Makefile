GO111MODULES    :=on
APP             :=summer
REGISTRY?       :=gcr.io/images
COMMIT_SHA      :=$(shell git rev-parse --short HEAD)
CMD_DIR         :=cmd
BIN_DIR         :=bin

#ci: clean vet build test cov
ci: clean build

.PHONY: build
## build: build the application
build: download
	@echo "Building..."
	CGO_ENABLED=0 GOOS=linux go build -o ${BIN_DIR}/${APP} $(CMD_DIR)/main.go

.PHONY: run
## run: runs go run main.go
run:
	go run -race $(CMD_DIR)/main.go

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning"
	@go clean

.PHONY: test
## test: runs go test with default values
test:
	go test -v -count=1 -race ./...

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

# helper rule for deployment
check-environment:
ifndef ENV
	$(error ENV not set, allowed values - `staging` or `production`)
endif

.PHONY: docker-build
## docker-build: builds the summer docker image to registry
docker-build: build
	docker build -t ${APP}:${COMMIT_SHA} .

.PHONY: docker-run
## docker-run: run the summer docker container
docker-run: docker-build
	docker run -p8080:3000 -d ${APP}:${COMMIT_SHA}

.PHONY: docker-push
## docker-push: pushes the summer docker image to registry
docker-push: check-environment docker-build
	docker push ${REGISTRY}/${ENV}/${APP}:${COMMIT_SHA}

.PHONY: help
## help: Prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
