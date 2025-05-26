install:
	go get -u github.com/bufbuild/buf/cmd/buf
	go install github.com/bufbuild/buf/cmd/buf

	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go

gen:
	buf generate

go-path:
	echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.bashrc