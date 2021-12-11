# auth-service

Service responsible for authentication/authorization

## Development

First, you need to install protobuf
```shell
brew install protobuf
```
Also, you need some tools like protoc-gen-go which you can install with
```shell
make install
```
Mocks and grpc definitions can be generated with
```shell
make generate
```