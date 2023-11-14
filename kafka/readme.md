### Stream logs from Kafka

Streams Kafka topics to ServiceNow Cloud Observability logs

#### Run

1) Download and unzip https://www.confluent.io/hub/confluentinc/kafka-connect-elasticsearch into `connect-plugins/`

2) Run Kafka

```
    docker-compose up

    # wait a few minutes...
```

3) Configure connector

```
    curl -X POST http://localhost:8083/connectors -H 'Content-Type: application/json' -d @plugin-config.json
```

4) Send some data

```
docker exec -i kafka bash -c "echo '{\"userId\": \"1\", \"action\": \"login\"}' | /usr/bin/kafka-console-producer --broker-list kafka:9092 --topic example-topic"

```