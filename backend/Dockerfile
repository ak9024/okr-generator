FROM golang:alpine AS builder

ARG HOST
ARG PORT
ARG ENV
ARG VERSION
ARG TOKEN

ENV HOST=${HOST}
ENV PORT=${PORT}
ENV ENV=${ENV}
ENV VERSION=${VERSION}
ENV TOKEN=${TOKEN}

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o backend
RUN ./backend env-generator

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/backend .
COPY --from=builder /app/.config.generated.toml .config.toml

ENTRYPOINT [ "./backend", "server" ]
