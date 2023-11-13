DOCKER_RELEASE_TAG ?= latest
APPLICATION_NAME ?=fermina-care-app
BUILD_NAME ?=fermina-care

dev:
	@go run cmd/main.go
build-app:
	@go build -o main-go cmd/main.go
tidy:
	@go mod tidy
d_build:
	docker build -t agungbhaskara/${BUILD_NAME}:${DOCKER_RELEASE_TAG} .
d_run_container:
	docker run -d -p 8080:8080 --name ${APPLICATION_NAME} agungbhaskara/${BUILD_NAME}:${DOCKER_RELEASE_TAG}