.PHONY: protos

protos:
	protoc --proto_path=proto --go_out=plugins=grpc:proto service.proto

install:
	go mod download

build:
	go build -o /balancer

run:
	./balancer