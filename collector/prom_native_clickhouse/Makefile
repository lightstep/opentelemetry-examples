DOCKER_NETWORK = clickhouse_integrations

run-client:
	docker run -it --rm --network ${DOCKER_NETWORK} --link clickhouse-server:clickhouse-server yandex/clickhouse-client --host clickhouse-server
