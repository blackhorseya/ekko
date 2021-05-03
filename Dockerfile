# build stage
FROM golang:alpine AS builder

WORKDIR /src

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
RUN go build -o app ./cmd/app

# final stage
FROM alpine:3

LABEL maintainer.name="blackhorseya"
LABEL maintainer.email="blackhorseya@gmail.com"

WORKDIR /app

COPY --from=builder /src/app ./

ENTRYPOINT ./app
