install_tools:
	@go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	@apt-get update
	@apt-get install -y protoc-gen-go
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

gen-grpc:
	@mkdir -p pkg/rpc
	@protoc \
		--go_opt=paths=source_relative --go_out=pkg/rpc \
		--go-grpc_opt=paths=source_relative --go-grpc_out=pkg/rpc \
		proto/merge.proto

clean-grpc:
	@rm -rf pkg/rpc