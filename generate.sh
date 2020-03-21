protoc -I/usr/local/include -I. -I$GOPATH/src  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. src/interfaces/rpc/pictweetpb/pictweet.proto
