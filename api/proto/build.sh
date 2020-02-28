protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=.  info/info.proto
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=.  user/user.proto
