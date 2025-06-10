#!/bin/bash
SCRIPT_PATH=$(realpath "$0")
SCRIPT_DIR=$(dirname "$SCRIPT_PATH")
SERVER_DIR=$SCRIPT_DIR/../api-docs
cd SERVER_DIR 
go get hangouts/gen
go mod tidy
