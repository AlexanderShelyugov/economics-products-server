# ECONOMICS PRODUCTS SERVER

This repository contains the source code for the products server of the Economics project.

You can build and run it locally or use the complete setup using [economics-ci](https://github.com/AlexanderShelyugov/economics-ci) repository.

# Table of contents
1. [Build](#Build)
    1. [Project](#Project)
    2. [Docker](#Docker)
2. [Run](#Run)
3. [Libraries](#Libraries)

## Build

### Project
In order to build the server run in shell
```
go build ./cmd/server.go
```

### Docker
You have two options to build a Docker image.
* If you have built a project according to [the previous step](#Project), you can just run 
```
docker build -f Dockerfile.lite -t economics:server-products
```
* Or if you don't want to setup **go**, build etc. you can just create a production build using
```
build.sh
```
or
```
docker build . -t economics:server-products
```

## Run

Same here:
```
./server
```

or
```
go run ./cmd/server.go
```
or
```
docker run -d --rm -p 80:80 economics:server-products
```
or
```
run.sh
```

## Libraries
This app uses [net/http](https://golang.org/pkg/net/http/) for running the HTTP server, [GORM](https://gorm.io/) as an ORM system.