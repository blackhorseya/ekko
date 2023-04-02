# build stage
FROM golang:alpine AS builder

ARG MAIN_PKG

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
COPY api ./api
RUN go build -o app ${MAIN_PKG}

# final stage
FROM alpine:3

LABEL maintainer.name="blackhorseya"
LABEL maintainer.email="blackhorseya@gmail.com"

WORKDIR /app

COPY --from=builder /src/app ./

ENTRYPOINT ./app
