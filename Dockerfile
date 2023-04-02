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

# build frontend
FROM node:alpine AS builder-f2e

WORKDIR /src

ENV NODE_OPTIONS=--openssl-legacy-provider
COPY web/package.json ./
RUN yarn install

COPY web/public ./public
COPY web/src ./src
RUN yarn build

# final stage
FROM alpine:3

LABEL maintainer.name="blackhorseya"
LABEL maintainer.email="blackhorseya@gmail.com"

WORKDIR /app

COPY --from=builder /src/app ./
COPY --from=builder-f2e /src/build ./web/build

ENTRYPOINT ./app
