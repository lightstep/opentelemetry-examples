FROM openjdk:18.0.1.1-jdk

RUN mkdir /app
WORKDIR /app
VOLUME /app

ENV JAR_TARGET "micrometer-demo-0.0.1-SNAPSHOT.jar"
ENTRYPOINT ["sh","-c","java -jar ./build/libs/${JAR_TARGET}"]