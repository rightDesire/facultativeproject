//go:generate bash -c "protoc -I . --go_out=paths=source_relative:./generated --go-grpc_out=paths=source_relative:./generated *.proto"
package proto
