FROM golang:1.23-alpine AS build-stage

WORKDIR /app

COPY ../go.mod ../go.sum ./

RUN go mod tidy

COPY ../ ./

RUN go build -o 01-api ./cmd

FROM alpine:latest AS production-stage

WORKDIR /app

COPY --from=build-stage /app/01-api /app/01-api

EXPOSE 8080

CMD ["./01-api"]
