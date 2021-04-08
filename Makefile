all: build

# 生成 application/api/rpc/activity.pb.go
proto: application/api/rpc/activity.proto
    protoc --go_out=plugins=grpc:./ application/api/rpc/activity.proto

build:
	go build -o bin/Seckill main.go

clean:
	rm bin/Seckill

.PHONY: clean build all