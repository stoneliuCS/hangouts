#!/bin/bash
SCRIPT_PATH=$(realpath "$0")
SCRIPT_DIR=$(dirname "$SCRIPT_PATH")
SERVER_DIR=$SCRIPT_DIR/../
cd $SERVER_DIR 
go run ./cmd/server/main.go
