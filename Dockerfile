FROM golang:1.24

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

ARG APP_NAME
ARG HTTP_PORT
ENV APP_NAME=${APP_NAME}
ENV HTTP_PORT=${HTTP_PORT}


RUN go build -o ${APP_NAME} ./cmd/app

EXPOSE ${HTTP_PORT}
