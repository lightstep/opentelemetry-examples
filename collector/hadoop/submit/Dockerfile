FROM curlimages/curl:7.82.0 as hadoop_curl
ARG JMX_JAR_VERSION=v1.14.0
USER root

RUN curl -L \
    --output /hadoop-mapreduce-examples-2.7.1-sources.jar \
    "https://repo1.maven.org/maven2/org/apache/hadoop/hadoop-mapreduce-examples/2.7.1/hadoop-mapreduce-examples-2.7.1-sources.jar"

FROM bde2020/hadoop-base:2.0.0-hadoop3.2.1-java8

COPY --from=hadoop_curl /hadoop-mapreduce-examples-2.7.1-sources.jar /opt/hadoop/applications/hadoop-mapreduce-examples-2.7.1-sources.jar

ENV JAR_FILEPATH="/opt/hadoop/applications/hadoop-mapreduce-examples-2.7.1-sources.jar"
ENV CLASS_TO_RUN="org.apache.hadoop.examples.WordCount"
ENV PARAMS="/input /output"

ADD run.sh /run.sh
RUN chmod a+x /run.sh

CMD ["/run.sh"]