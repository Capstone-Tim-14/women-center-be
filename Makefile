DOCKER_RELEASE_TAG ?= latest
APPLICATION_NAME ?=fermina-care-app
BUILD_NAME ?=fermina-care
DOCKER_USERNAME ?=capstone14

dev:
	@go run cmd/main.go
build-app:
	@go build -o main-go cmd/main.go
tidy:
	@go mod tidy
d_stop:
	docker stop ${APPLICATION_NAME} || true
d_rename:
	docker rm ${APPLICATION_NAME} || true
d_build:
	docker build -t ${DOCKER_USERNAME}/${BUILD_NAME}:${DOCKER_RELEASE_TAG} .
d_run_container:
	docker run -d -p 8080:8080 --name ${APPLICATION_NAME} ${DOCKER_USERNAME}/${BUILD_NAME}:${DOCKER_RELEASE_TAG}
d_push:
	docker push ${DOCKER_USERNAME}/${BUILD_NAME}:${DOCKER_RELEASE_TAG}
d_pull:
	docker pull ${DOCKER_USERNAME}/${BUILD_NAME}:${DOCKER_RELEASE_TAG}
d_run_ngrok:
	docker run -it -e NGROK_AUTHTOKEN=2XJ401tq0VyCnCf66NlRrf9Mgih_5vNZpQ6n5PZk3SVoithRh -p 8080:8080 ngrok/ngrok http 35.223.26.211:8080