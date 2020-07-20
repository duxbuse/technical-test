ARG VERSION=test
ARG COMMITSHA=abc123

FROM golang:1.14-alpine as builder
    WORKDIR /app

    RUN apk update && apk add --no-cache git

    COPY go.mod go.sum /app/
    RUN go mod download
    RUN go mod verify

    COPY main.go /app/
    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/test2

FROM scratch
    ENV VERSION = $VERSION
    ENV COMMITSHA = $COMMITSHA
    COPY --from=builder /app/bin/test2 /test2
    ENTRYPOINT ["/test2"]
    EXPOSE 8000