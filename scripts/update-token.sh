#!/bin/bash

cp example.env .env
sed -i '' "s/<ACCESS TOKEN>/${TOKEN}/" .env
cp example-collector-config.yaml ./collector/collector-config.yaml
sed -i '' "s/<ACCESS TOKEN>/${TOKEN}/" ./collector/collector-config.yaml