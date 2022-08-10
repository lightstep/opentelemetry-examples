#!/bin/bash

VAULT_RETRIES=5
echo "Vault is starting..."
until vault status > /dev/null 2>&1 || [ "$VAULT_RETRIES" -eq 0 ]; do
	echo "Waiting for vault to start...: $((VAULT_RETRIES--))"
	sleep 1
done

echo "Authenticating to vault..."
vault login token=vault-plaintext-root-token

echo "Confirm the login was successful..."
vault token lookup | grep policies

echo "Define a prometheus-metrics ACL policy that grants read capabilities to the metrics endpoint..."
vault policy write prometheus-metrics - << EOF
path "/sys/metrics" {
  capabilities = ["read"]
}
EOF

# echo "Write the token ID to prometheus-token file to access the Vault telemetry metrics endpoint..."
# vault token create \
#   -field=token \
#   -policy prometheus-metrics \
#   > prometheus-config/prometheus-token

echo "Setting vault..."
vault secrets enable -version=2 -path=my.secrets kv

echo "Adding entries..."
vault kv put my.secrets/dev username=test_user
vault kv put my.secrets/dev password=test_password

echo "Complete..."