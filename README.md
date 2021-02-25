# go-grpc-beautiful-architecture


## Why gRPC?

gRPC is a modern open source high performance RPC framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.


## Communication flow


![myimage-alt-tag](https://cdn-images-1.medium.com/max/1600/1*4COSgdUTh2fmJJJpONJVVw.png) 


## Dependencies

### GO version

- [GO v1.15](https://golang.org)

### Base Framework
- [GRPC](https://grpc.io/docs/languages/go/quickstart/)

### DB Connection
- [Gorm](https://github.com/jinzhu/gorm)

### Other
- [Configor v1.2.0](https://github.com/jinzhu/configor)


## Exec

### Reder Protocol Buffers
```sh
$ git clone https://github.com/javierlecca/go-grpc-beautiful-architecture.git
$ sh renderProto.sh
```
### Run Project
```sh
$ go run server.go
```