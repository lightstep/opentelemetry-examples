#!/bin/bash

docker-compose up config-generator

if ! git diff-index --quiet HEAD --; then
    echo "config has changed"
    git diff
    exit 1
fi
