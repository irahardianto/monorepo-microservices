#build stage
FROM golang:1.11.4 AS builder
WORKDIR $GOPATH/src/github.com/irahardianto/monorepo-microservices
COPY go.mod .
COPY go.sum .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go mod download
COPY . .
RUN go build -a -installsuffix cgo -o goapp cmd/bookings/main.go

#final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/irahardianto/monorepo-microservices/goapp .
CMD ["./goapp"]

EXPOSE 8080