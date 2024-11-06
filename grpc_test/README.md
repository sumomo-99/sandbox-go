## gRPCテストサーバー

## gRPCコードの生成
```console
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
grpc_test/infra.proto
```

## gRPCサーバーの起動
```console
# 80ポートで起動
go run main.go

# 8080ポートで起動
go run main.go --port 8080
```
