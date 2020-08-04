default: build

build-client:
	mvn -f pom-client.xml package

build-server:
	mvn -f pom-server.xml package

build: build-client build-server

run-client:
	java -jar target/lightstep-otlp-client.jar

run-server:
	java -jar target/lightstep-otlp-server.jar

clean:
	mvn -f pom-client.xml clean