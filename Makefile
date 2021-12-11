install-tools:
	@echo "installing latest versions of protoc-gen-go and protoc-gen-go-grpc..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate:
	@echo "generating mocks and models..."
	@go generate ./...