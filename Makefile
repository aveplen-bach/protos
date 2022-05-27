.PHONY: s3file
s3file:
	mkdir -p s3file && \
	protoc -I=. \
    --go_out s3file \
    --go_opt paths=source_relative \
    --go-grpc_out s3file \
    --go-grpc_opt paths=source_relative \
    s3file.proto
