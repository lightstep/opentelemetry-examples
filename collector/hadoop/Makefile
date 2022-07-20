DOCKER_NETWORK = hadoop_integrations
ENV_FILE = hadoop.env
tag := 2.0.0-hadoop3.2.1-java8

wordcount:
	docker build -t hadoop-wordcount ./submit
	docker run --network ${DOCKER_NETWORK} --env-file ${ENV_FILE} bde2020/hadoop-namenode:$(tag) hdfs dfs -mkdir -p /input/
	docker run --network ${DOCKER_NETWORK} --env-file ${ENV_FILE} bde2020/hadoop-namenode:$(tag) hdfs dfs -copyFromLocal -f /opt/hadoop-3.2.1/README.txt /input/
	docker run --network ${DOCKER_NETWORK} --env-file ${ENV_FILE} hadoop-wordcount
	docker run --network ${DOCKER_NETWORK} --env-file ${ENV_FILE} bde2020/hadoop-namenode:$(tag) hdfs dfs -cat /output/*
	docker run --network ${DOCKER_NETWORK} --env-file ${ENV_FILE} bde2020/hadoop-namenode:$(tag) hdfs dfs -rm -r /output
	docker run --network ${DOCKER_NETWORK} --env-file ${ENV_FILE} bde2020/hadoop-namenode:$(tag) hdfs dfs -rm -r /input

remove_all_images:
	docker rmi $(shell docker images 'bde2020/hadoop-namenode' -q)
	docker rmi $(shell docker images 'bde2020/hadoop-datanode' -q)
	docker rmi $(shell docker images 'bde2020/hadoop-resourcemanager' -q)
	docker rmi $(shell docker images 'bde2020/hadoop-historyserver' -q)
	docker rmi $(shell docker images 'bde2020/hadoop-nodemanager' -q)
	docker rmi $(shell docker images 'hadoop_otel-collector' -q)

