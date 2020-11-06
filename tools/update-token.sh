#!/bin/bash

cp ./config/example.env .env
sed -i'' "s#<ACCESS TOKEN>#${TOKEN}#" .env
sed -i'' "s#<ORG_NAME>#${ORG_NAME}#" .env
sed -i'' "s#<PROJECT_NAME>#${PROJECT_NAME}#" .env
sed -i'' "s#<API_KEY>#${API_KEY}#" .env
cp ./config/example-collector-config.yaml ./config/collector-config.yaml
sed -i'' "s#<ACCESS TOKEN>#${TOKEN}#" ./config/collector-config.yaml
