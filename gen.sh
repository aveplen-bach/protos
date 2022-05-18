#!/bin/sh

protoc -I=. \
    --go_out . \
    --go_opt paths=source_relative \
    --go-grpc_out . \
    --go-grpc_opt paths=source_relative \
    facerec.proto \
    s3g.proto

./venv/bin/python -m grpc_tools.protoc -I=.\
    --python_out . \
    --grpc_python_out . \
    facerec.proto \
    s3g.proto