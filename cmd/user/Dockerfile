ARG DOCKER_ID

FROM golang:1.16-alpine AS builder

# Download dependencies
RUN mkdir -p ${GOPATH}/src/app/cmd/company && \
    mkdir -p ${GOPATH}/src/app/cmd/gateway && \
    mkdir -p ${GOPATH}/src/app/cmd/user
WORKDIR ${GOPATH}/src/app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify
# Build binary
COPY ./cmd/user ./cmd/user
COPY ./lib ./lib
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o main ./cmd/user

# Build minimum image
FROM alpine:3.14 as production

# Copy our static executable.
WORKDIR /app
COPY --from=builder /go/src/app/main .
# Run application
ENTRYPOINT ["/app/main"]

# Expose default port
EXPOSE 9000
