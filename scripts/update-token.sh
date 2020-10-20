#!/bin/bash

cp example.env .env
sed -i'' "s#<ACCESS TOKEN>#${TOKEN}#" .env
sed -i'' "s#<ORG_NAME>#${ORG_NAME}#" .env
sed -i'' "s#<PROJECT_NAME>#${PROJECT_NAME}#" .env
sed -i'' "s#<API_KEY>#${API_KEY}#" .env
cp example-collector-config.yaml ./collector/collector-config.yaml
sed -i'' "s#<ACCESS TOKEN>#${TOKEN}#" ./collector/collector-config.yaml
