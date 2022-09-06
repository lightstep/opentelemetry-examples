#!/bin/bash

cp ./config/example.env .env
sed -i.bak "s#<ACCESS TOKEN>#${TOKEN}#" .env
sed -i.bak "s#<ORG_NAME>#${ORG_NAME}#" .env
sed -i.bak "s#<PROJECT_NAME>#${PROJECT_NAME}#" .env
sed -i.bak "s#<API_KEY>#${API_KEY}#" .env
cp ./config/example-collector-config.yaml ./config/collector-config.yaml
sed -i.bak "s#<ACCESS TOKEN>#${TOKEN}#" ./config/collector-config.yaml
