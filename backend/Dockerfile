# syntax = docker/dockerfile:1.0-experimental
FROM golang:1.17 AS build
WORKDIR /app
COPY . .
RUN --mount=type=cache,target=/root/.cache go build -mod vendor -ldflags '-w -s' -o /app/lol -v ./cmd/main.go

FROM debian:latest
WORKDIR /app
COPY  --from=build /app/lol .
COPY --from=build /app/auth-public.pem .
COPY --from=build /app/auth-private.pem .
ENTRYPOINT ["/app/lol"]
