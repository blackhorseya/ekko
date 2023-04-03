# ekko

[![Build](https://github.com/blackhorseya/ekko/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/blackhorseya/ekko/actions/workflows/build.yml)
[![codecov](https://codecov.io/gh/blackhorseya/ekko/branch/main/graph/badge.svg?token=SV4V6G6QZJ)](https://codecov.io/gh/blackhorseya/ekko)
[![Go Report Card](https://goreportcard.com/badge/github.com/blackhorseya/ekko)](https://goreportcard.com/report/github.com/blackhorseya/ekko)
[![Go Reference](https://pkg.go.dev/badge/github.com/blackhorseya/ekko)](https://pkg.go.dev/github.com/blackhorseya/ekko)
[![Release](https://img.shields.io/github/release/blackhorseya/ekko)](https://github.com/blackhorseya/ekko/releases/latest)
[![GitHub license](https://img.shields.io/github/license/blackhorseya/ekko)](https://github.com/blackhorseya/ekko/blob/main/LICENSE)

The main purpose of this project is to practice `golang` and design the system architecture

## Live Preview

[Demo](https://ekko.seancheng.space)

## Concepts

### Project layout

The project layout is designed with reference to the [Standard Go Project Layout](https://github.com/golang-standards/project-layout) combined with the concept of [monorepo](https://monorepo.tools/)

### System design

Design the corresponding microservice based on Domain-Driven Design

**Domain**

- `Account`: define a user
- `Task`: as a representation of all tickets

## Functional

### Main

- add tasks
- list all tasks by completed
- remove tasks
- modify tasks

### Extra

- task with tags field
- task with startAt and endAt field

## Technical

### Languages

- Go
- React

### Toolchain

- [Gin](https://github.com/gin-gonic/gin)
- [Protobuf](https://developers.google.com/protocol-buffers)
- [Wire](https://github.com/google/wire)
- [Testify](https://github.com/stretchr/testify)
- [Mockgen](https://github.com/golang/mock)

### Infrastructures

- Mariadb
- Docker
- GCR
- GKE
- Helm 3
- Terraform

### CI/CD

- GitHub Action
- [Codecov](https://codecov.io/)
