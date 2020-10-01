# Fetcher API

This is simple fetcher API, which provides service for saving URLs and getting data from it with specified time interval.

## How to build
```
$ go generate
$ go build .
```

## How to run

#### Normal way
```
$ go run .
```
#### With docker
```
$ docker-compose up
```

## Configuration


## Project structure

- db - contains database schema
- engine - contains source code of saving url's and fetching mechanism
- grpc - contains gRPC API server with gRPC gateway
- rest - contains HTTP API server with Chi router

## Description

Some info about how this project was designed and why that way and how to use them.