#!/bin/bash

SRC_DIR=$1
DOCS_URL=$2

find $SRC_DIR -type f -name "*.go" -exec sed -i \
    "s|localhost:8080|${DOCS_URL}|g" {} +
