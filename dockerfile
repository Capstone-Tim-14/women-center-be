FROM golang:alpine3.18 as go-app

RUN apk add --no-cache make

RUN mkdir /app
WORKDIR /app
ADD . .
RUN make tidy
RUN make build-app

FROM alpine:3.18
WORKDIR /root/
COPY --from=go-app /app/main-go .
EXPOSE 8080
CMD [ "./main-go" ]