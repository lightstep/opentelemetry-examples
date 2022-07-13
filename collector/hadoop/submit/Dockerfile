FROM bde2020/hadoop-base:2.0.0-hadoop3.2.1-java8

COPY hadoop-mapreduce-examples-2.7.1-sources.jar /opt/hadoop/applications/hadoop-mapreduce-examples-2.7.1-sources.jar

ENV JAR_FILEPATH="/opt/hadoop/applications/hadoop-mapreduce-examples-2.7.1-sources.jar"
ENV CLASS_TO_RUN="org.apache.hadoop.examples.WordCount"
ENV PARAMS="/input /output"

ADD run.sh /run.sh
RUN chmod a+x /run.sh

CMD ["/run.sh"]