FROM golang:1.14-alpine as builder
    WORKDIR /app

    RUN apk update --no-cache && apk add --no-cache git

    COPY go.mod go.sum /app/
    RUN go mod download
    RUN go mod verify

    COPY main.go /app/
    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/golang-test

FROM scratch
    COPY --from=builder /app/golang-test /app/golang-test
    ENTRYPOINT ["/app/golang-test"]
    EXPOSE 8000