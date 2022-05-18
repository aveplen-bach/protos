#!/bin/sh

rm -rf facerec
mkdir facerec

protoc -I=. \
    --go_out facerec \
    --go_opt paths=source_relative \
    --go-grpc_out facerec \
    --go-grpc_opt paths=source_relative \
    facerec.proto

./venv/bin/python -m grpc_tools.protoc -I=.\
    --python_out facerec \
    --grpc_python_out facerec \
    facerec.proto


rm -rf s3g
mkdir s3g

protoc -I=. \
    --go_out s3g \
    --go_opt paths=source_relative \
    --go-grpc_out s3g \
    --go-grpc_opt paths=source_relative \
    s3g.proto

./venv/bin/python -m grpc_tools.protoc -I=.\
    --python_out s3g \
    --grpc_python_out s3g \
    s3g.proto
