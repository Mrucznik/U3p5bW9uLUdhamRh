# Fetcher API

This is the simple fetcher API, which provides service for saving URLs and getting data from it with a specified time interval.

## Endpoints

This API provides endpoints:
- GET /api/fetcher - Gets all stored URLs.
- POST /api/fetcher - Create a URL and runs worker, that fetch data from url by specified interval.
- DELETE /api/fetcher/{id} - Delete a URL and stops worker.
- GET /api/fetcher/{id}/history - Get URL fetching history.

## Documentation
You can find swagger-ui documentation [here](https://mrucznik.github.io/url-fetcher).

## How to build
You need the Google protocol buffers compiler to build this project. You can download it from the official repository: https://github.com/protocolbuffers/protobuf/releases 
```
$ go install \
      github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
      github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
      github.com/golang/protobuf/protoc-gen-go
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

You can configure this project by setting up environment variables:
- USE_GRPC - true/false. Should we serve API by gRPC or by Chi.
- PORT - port on which HTTP server will be listening on.
- GRPC_PORT - port on which gRPC server will be listening on.
- USE_DATABASE - true/false. Should service use in memory implementation or database.
- DSN - database connection string.

## Project structure

- db - contains database schema and database connection configuration.
- docs - documentation files (swagger file and swagger-ui interface)
- engine - contains source code of saving URLs and fetching mechanism.
    - database - implementation with storing URLs in a MySQL database.
    - in_memory - implementation with storing URLs in memory.
- grpc - contains gRPC API server with gRPC gateway.
    - proto - protobuf files.
    - third_party - third party files.
- rest - contains HTTP API server with Chi router. Not implemented.
