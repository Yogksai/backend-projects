FROM golang:1.23.0-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /docker-gs-ping

FROM alpine:latest

COPY --from=build /docker-gs-ping /docker-gs-ping

ENTRYPOINT [ "/docker-gs-ping" ]
