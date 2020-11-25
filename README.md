# RPCdemo
This project demonstrate how works comunication between client and server with gRPC, doing a Nine Table Calc.

To run the project, you would just run the commands below:

Run Server:
```
go run server/main.go
```
*The server will be listen to port 8080 by default, remember to close it after you done.*

Run Client:
```
go run client/main.go
```
*The Client will send a Procedure Call to server and it will die, you can edit to your linking.*

To edit and play with the code, you will need to follow some more steps:

Install the protoc compiler that is used to generate gRPC service code. 
```
 go get -u github.com/golang/protobuf/protoc-gen-go
```
The compiler plugin, protoc-gen-go, will be installed in $GOBIN, defaulting to $GOPATH/bin. It must be in your $PATH for the protocol compiler, protoc, to find it.
```
 export PATH=$PATH:$GOPATH/bin
```

Create client and service interfaces
```
protoc --go_out=plugins=grpc:. ./proto/ninetable.proto 
```

You can find more about **gRPC** in golang [here](https://grpc.io/docs/languages/go/).