#!/bin/bash

#生成go文件
echo "golang proto build"
protoc --go_out=plugins=grpc:. ./dc_rpc.proto