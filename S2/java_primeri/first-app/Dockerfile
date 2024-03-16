FROM openjdk:8-jdk-alpine
LABEL maintainer="danijelradakovic@uns.ac.rs"
EXPOSE 8080
WORKDIR /app
COPY ./servers.jar ./
CMD ["java", "-jar", "servers.jar"]
