
FROM golang:1.14-alpine as builder
    ARG VERSION="test"
    ARG COMMITSHA="abc123"

    WORKDIR /app

    RUN apk update && apk add --no-cache git

    COPY go.mod go.sum /app/
    RUN go mod download
    RUN go mod verify

    COPY main.go /app/
    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
        -ldflags "-X main.VersionNumber=$VERSION -X main.LastCommitSha=$COMMITSHA" \
        -o /app/bin/test2

FROM scratch
    COPY --from=builder /app/bin/test2 /test2
    ENTRYPOINT ["/test2"]
    EXPOSE 8000