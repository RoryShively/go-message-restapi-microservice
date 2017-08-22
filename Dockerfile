# Run golang:alpine as build stage
FROM golang:alpine  AS build-env
RUN apk add --no-cache git mercurial
WORKDIR /go/src/app
ADD . /go/src/app/
RUN go get && go build -o message-service

# Final image
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/app/message-service ./message-service
ENTRYPOINT ./message-service
EXPOSE 8080
