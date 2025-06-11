#!/bin/bash


SCRIPT_PATH=$(realpath "$0")
SCRIPT_DIR=$(dirname "$SCRIPT_PATH")

OAS_SCRIPT_DIR=$SCRIPT_DIR/../api-docs
cd $OAS_SCRIPT_DIR
bun run index.ts
