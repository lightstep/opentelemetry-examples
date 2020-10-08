#!/bin/bash

docker-compose up config-generator

if [ -z $(git diff --quiet --exit-code ./integration) ]; then 
    echo "config has changed"
    git status
    exit 1
fi
