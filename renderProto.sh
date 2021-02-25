export PATH=$PATH:$HOME/go/bin
export PATH=$PATH:/usr/local/go/bin
protoc -I proto --go_out=plugins=grpc:internal ./proto/*.proto