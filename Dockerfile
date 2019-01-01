#build stage
FROM golang:1.11.4 AS builder
ARG SERVICE_NAME
WORKDIR $GOPATH/src/github.com/irahardianto/monorepo-microservices
COPY go.mod .
COPY go.sum .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go mod download
COPY . .
RUN go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o goapp cmd/$SERVICE_NAME/main.go

#final stage
FROM debian:9.6-slim
WORKDIR /root/
RUN mkdir -p ./cmd/bookings
COPY --from=builder /go/src/github.com/irahardianto/monorepo-microservices/goapp .
COPY --from=builder /go/src/github.com/irahardianto/monorepo-microservices/config/config.yaml ./config/
CMD ["./goapp"]

EXPOSE 8080