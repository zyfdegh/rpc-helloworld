default: help

help:
	@echo -e "protoc: \n\t Download and install protoc v3.3.0"
	@echo -e "gen: \n\t Generate go source from proto file"
	@echo -e "dep-update: \n\t Sync dependencies to vendor/"
	@echo -e "fmt: \n\t Format go source codes"
	@echo -e "run-server: \n\t Build and run server"
	@echo -e "run-client: \n\t Build and run client"
	@echo -e "help: \n\t Display this help"

dep-init:
	go get github.com/kardianos/govendor
	govendor init

dep-update:
	govendor sync
	govendor remove +unused
	govendor add +external

fmt:
	go fmt $(go list ./... | grep -v /vendor/)

protoc:
	wget -o /tmp/protoc.zip https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip
	unzip -p /tmp/protoc.zip bin/protoc > /usr/local/bin/protoc
	chmod +x /usr/local/bin/protoc
	rm /tmp/protoc.zip

run-server:
	go run greeter_server/main.go

run-client:
	go run greeter_client/main.go

gen:
	protoc --go_out=plugins=grpc:. helloworld/helloworld.proto
