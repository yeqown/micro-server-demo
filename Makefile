gen-proto:
	protoc -I./api/proto foo.proto --go_out=plugins=grpc:api/protogen

build:
	go build -o app ./cmd

run-dev: build
	ENV=dev ./app start -rpcPort=8080