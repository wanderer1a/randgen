FROM golang:1.16-alpine as builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /main main.go
RUN ls -la /
FROM alpine:3
COPY --from=builder main /bin/main
WORKDIR /bin
ENTRYPOINT ["/bin/main"]