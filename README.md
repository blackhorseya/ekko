# Todo List

[![Build](https://github.com/blackhorseya/todo-app/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/blackhorseya/todo-app/actions/workflows/build.yml)
[![codecov](https://codecov.io/gh/blackhorseya/todo-app/branch/main/graph/badge.svg?token=SV4V6G6QZJ)](https://codecov.io/gh/blackhorseya/todo-app)
[![Go Report Card](https://goreportcard.com/badge/github.com/blackhorseya/todo-app)](https://goreportcard.com/report/github.com/blackhorseya/todo-app)
[![Go Reference](https://pkg.go.dev/badge/github.com/blackhorseya/todo-app)](https://pkg.go.dev/github.com/blackhorseya/todo-app)
[![Release](https://img.shields.io/github/release/blackhorseya/todo-app)](https://github.com/blackhorseya/todo-app/releases/latest)
[![GitHub license](https://img.shields.io/github/license/blackhorseya/todo-app)](https://github.com/blackhorseya/todo-app/blob/main/LICENSE)

The main purpose of this project is to practice `golang` and design the system architecture

## Live Preview

[Demo](https://todo.seancheng.space)

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
- [Mockery](https://github.com/vektra/mockery)

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
