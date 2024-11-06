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

## Docker
**イメージのビルド**
```console
docker build -t grpc-test-server .
```
**コンテナの起動**
```console
# 80ポートで起動
docker run -d --rm -p 80:80 grpc-test-server
# 8080ポートで起動
docker run -d --rm -p 8080:8080 grpc-test-server --port 8080
```
