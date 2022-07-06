FROM jboss/wildfly

ARG BOOTCLASS_PATH="-Xbootclasspath/a:"

RUN /opt/jboss/wildfly/bin/add-user.sh admin admin --silent

ARG JAVA_OPTS_1
ARG LOGMANAGER
ARG LOGMANAGER_JAR
ARG WILDLFY_COMMON_JAR

ENV JAVA_OPTS="${JAVA_OPTS}${JAVA_OPTS_1}${LOGMANAGER}${BOOTCLASS_PATH}${JBOSS_HOME}${LOGMANAGER_JAR}${JBOSS_HOME}${WILDLFY_COMMON_JAR}"


ENTRYPOINT /opt/jboss/wildfly/bin/standalone.sh \  
    -b=0.0.0.0 \
    -bmanagement=0.0.0.0 \
    -Djboss.server.default.config=standalone-full-ha.xml \
    -Djava.net.preferIPv4Stack=true
