FROM jboss/wildfly:9.0.2.Final

ARG BOOTCLASS_PATH="-Xbootclasspath/a:"

RUN /opt/jboss/wildfly/bin/add-user.sh admin admin --silent

ARG JAVA_OPTS
ARG LOGMANAGER
ARG LOGMANAGER_JAR
ARG WILDLFY_COMMON_JAR

USER root

ENV JAVA_OPTS="${JAVA_OPTS}${JAVA_OPTS}${LOGMANAGER}${BOOTCLASS_PATH}${JBOSS_HOME}${LOGMANAGER_JAR}${JBOSS_HOME}${WILDLFY_COMMON_JAR}"

USER jboss

# cleanup
RUN rm -rf $JBOSS_HOME/standalone/configuration/standalone_xml_history

# automatically launch the standalone configuration
CMD ["/opt/jboss/wildfly/bin/standalone.sh", "-b", "0.0.0.0", "-bmanagement", "0.0.0.0", "-c", "standalone-full.xml"]
# ENTRYPOINT /opt/jboss/wildfly/bin/standalone.sh -Djboss.bind.address.management=0.0.0.0 -Djboss.bind.address=0.0.0.0
