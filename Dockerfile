# Stage 0, build binary
FROM golang:1.16-alpine3.12 AS builder
WORKDIR /build

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags musl -o server cmd/server/server.go
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags musl -o client cmd/client/client.go

# Stage 1, build container with only necessary files
FROM alpine:3.12

WORKDIR /GoChat/

COPY --from=builder /build/server .
COPY --from=builder /build/client .

EXPOSE 9099

ENTRYPOINT ["/GoChat/server"]
