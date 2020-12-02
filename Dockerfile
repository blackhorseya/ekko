# build stage
FROM golang:alpine AS builder

LABEL app="todo"

WORKDIR /src

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd ./cmd
RUN go build -o app ./cmd/app

# final stage
FROM alpine

LABEL maintainer="blackhorseya"
LABEL mail="blackhorseya@gmail.com"
LABEL app="todo"

WORKDIR /app

COPY --from=builder /src/app ./

ENTRYPOINT ./app
