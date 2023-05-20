#build stage
FROM golang:1.20.4-bullseye

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
