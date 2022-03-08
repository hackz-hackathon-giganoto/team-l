#!/bin/sh

set -xe

SERVER_OUTPUT_DIR=grpc-server/messenger
CLIENT_OUTPUT_DIR=client/src/messenger

protoc --version
protoc --proto_path=proto messenger.proto \
  --go_out=plugins="grpc:${SERVER_OUTPUT_DIR}" \
  --js_out=import_style=commonjs:${CLIENT_OUTPUT_DIR} \
  --grpc-web_out=import_style=typescript,mode=grpcwebtext:${CLIENT_OUTPUT_DIR}

# protoc --proto_path=proto messenger.proto\
#     --js_out=import_style=commonjs:client/messenger \
#     --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:client/messenger \
#     --go-grpc_out=grpc-server/messenger \
#     --go_out=grpc-server/messenger