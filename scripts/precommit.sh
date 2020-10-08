#!/bin/bash

docker-compose up config-generator

if ! git diff --quiet; then
    echo;
    echo 'Working tree is not clean, did you forget to run "./scripts/precommit"?';
    echo;
    git status;
    exit 1;
fi