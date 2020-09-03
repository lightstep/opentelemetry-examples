.PHONY: build clean

specialagent_url := https://repo1.maven.org/maven2/io/opentracing/contrib/specialagent/opentracing-specialagent/1.7.0/opentracing-specialagent-1.7.0.jar

otel_auto_url := https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/latest/download/opentelemetry-javaagent-all.jar

build: opentracing-specialagent-1.7.0.jar
	mvn package

opentracing-specialagent-1.7.0.jar:
	wget ${specialagent_url}

build-otel: otel-latest
	mvn package

otel-latest:
	wget ${otel_auto_url}

clean:
	mvn clean
