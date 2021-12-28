.PHONY: install
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: generate_proto
generate_proto: install
	echo $(shell go env)
	protoc \
        -I=. \
        --go_out=. \
        --go_opt=module=github.com/wansnow/calculation_server \
        --go-grpc_out=. \
        --go-grpc_opt=module=github.com/wansnow/calculation_server \
        $(shell find proto -name '*.proto')

.PHONY: remove_tag
remove_tag:
	go run ./entrypoint/tool/remove_tag/main.go