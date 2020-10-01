# Build Stage
FROM golang:alpine AS build-stage

RUN apk update && apk add protoc

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

# Install plugins for protoc
RUN go install \
           github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
           github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
           github.com/golang/protobuf/protoc-gen-go

COPY . .

RUN go generate
RUN go build -o /build/app .


# Final Stage
FROM alpine:latest

WORKDIR /service

COPY --from=build-stage /build/app /service/app
RUN chmod +x /service/app

EXPOSE 8080

CMD ["/service/app"]
