# build stage
FROM golang:alpine AS builder

WORKDIR /src

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd ./cmd
RUN go build -o app ./cmd/app

# final stage
FROM alpine

WORKDIR /app

COPY --from=builder /src/app ./

ENTRYPOINT ./app
