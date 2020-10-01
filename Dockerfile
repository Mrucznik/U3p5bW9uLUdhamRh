# Build Stage
FROM golang:alpine AS build-stage

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

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
