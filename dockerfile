FROM golang:alpine3.18 as go-app

RUN apk add --no-cache make

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN make tidy
RUN make build-app

FROM alpine:3.18
WORKDIR /root/
COPY --from=go-app /app/main-go .
COPY env.yaml /root/env.yaml
EXPOSE 8080
CMD [ "./main-go" ]