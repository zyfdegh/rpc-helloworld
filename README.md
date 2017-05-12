# rpc-helloworld

# Summary
Example to show how to use RPC in Golang.

# Prerequisites
* Golang 1.6+
* make
* protoc 3.0+

# Clone
```sh
$ mkdir -p $GOPATH/src/github.com/zyfdegh/

$ git clone https://github.com/zyfdegh/rpc-helloworld \
	$GOPATH/src/github.com/zyfdegh/

$ cd $GOPATH/src/github.com/zyfdegh/
```

# Install protoc
```sh
$ make protoc
```

# Build & Run

Retrieve dependent packages,
```sh
$ make dep-update
```

Build and run,
```sh
$ make run-server
go run greeter_server/main.go
2017/05/12 17:06:20 server listenning on :50051
```

In a new terminal,
```sh
$ make run-client
go run greeter_client/main.go
2017/05/12 17:06:31 Greeting: Hello world
2017/05/12 17:06:31 Greeting: Fuck world
```

# Others
See
```sh
$ make help
```
