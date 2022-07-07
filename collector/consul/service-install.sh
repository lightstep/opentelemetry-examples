###
# Container 1
###

# Download sample data layer service, Counting service
docker exec -d consul-client1 wget https://github.com/hashicorp/demo-consul-101/releases/download/0.0.3.1/counting-service_linux_amd64.zip
sleep 7s
# Unzip Counting service
docker exec -d consul-client1 unzip counting-service_linux_amd64.zip
sleep 7s
# Start Counting service as background process in container
docker exec -d consul-client1 ./counting-service_linux_amd64 &
sleep 1s
# Start Consul Sidecar Proxy for Counting service
docker exec -d consul-client1 consul connect proxy -sidecar-for counting &
sleep 1s

###
# Container 2
###

# Download sample data layer service, Dashboard service
docker exec -d consul-client2 wget https://github.com/hashicorp/demo-consul-101/releases/download/0.0.3.1/dashboard-service_linux_amd64.zip
sleep 7s
# Unzip Counting service
docker exec -d consul-client2 unzip dashboard-service_linux_amd64.zip
sleep 7s
# Start Dashboard service as background process in container
docker exec -d consul-client2 ./dashboard-service_linux_amd64 &
sleep 1s
# Start Consul Sidecar Proxy for Dashboard service
docker exec -d consul-client2 consul connect proxy -sidecar-for dashboard &
sleep 1s

# Create Consul intention with Dashboard as source and Counting as destination
docker exec -d consul-client1 consul config write intention-config-entry.json
sleep 1s
