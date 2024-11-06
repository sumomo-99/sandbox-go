FROM golang:alpine AS builder

WORKDIR /app

RUN apk add --no-cache git protoc

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY . .

RUN protoc --go_out=. --go-grpc_out=. grpc_test/*.proto

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app


FROM alpine:latest

WORKDIR /app

COPY --from=builder /go/bin/app .

ENTRYPOINT ["./app"]
