protoc --proto_path=./proto --proto_path=./proto/libs/ \
    --go_out=./server/pb --go_opt paths=source_relative \
    --plugin=$GOPATH/bin/protoc-gen-grpc-gateway.exe \
    --plugin=$GOPATH/bin/protoc-gen-openapiv2.exe \
    --plugin=$GOPATH/bin/protoc-gen-go-grpc.exe \
    --go-grpc_out=./server/pb --go-grpc_opt paths=source_relative \
    --grpc-gateway_out=./server/pb \
    --grpc-gateway_opt allow_delete_body=true,logtostderr=true,paths=source_relative,repeated_path_param_separator=ssv \
    ./proto/api.proto

protoc --proto_path=./proto --proto_path=./proto/libs/ \
    --go_out=./server/pb --go_opt paths=source_relative \
    --plugin=$GOPATH/bin/protoc-gen-gorm.exe \
    --gorm_out=. \
    ./proto/gorm.proto 