FROM jboss/wildfly

ARG BOOTCLASS_PATH="-Xbootclasspath/a:"

RUN /opt/jboss/wildfly/bin/add-user.sh admin admin --silent

ARG JAVA_OPTS
ARG LOGMANAGER
ARG LOGMANAGER_JAR
ARG WILDLFY_COMMON_JAR

ENV JAVA_OPTS="${JAVA_OPTS}${JAVA_OPTS}${LOGMANAGER}${BOOTCLASS_PATH}${JBOSS_HOME}${LOGMANAGER_JAR}${JBOSS_HOME}${WILDLFY_COMMON_JAR}"

ENTRYPOINT /opt/jboss/wildfly/bin/standalone.sh -Djboss.bind.address.management=0.0.0.0 -Djboss.bind.address=0.0.0.0
