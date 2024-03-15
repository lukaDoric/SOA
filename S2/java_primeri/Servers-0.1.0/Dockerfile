FROM node:alpine3.11 AS frontBuild

LABEL maintainer="danijelradakovic@uns.ac.rs"

ARG REPOSITORY_OWNER=DanijelRadakovic
ARG APPLLICATION_REPOSITORY=Servers-Front
ARG APPLICATION_VERSION=v1.0.0

WORKDIR /usr/src
RUN apk --update --no-cache add curl tar && \
    curl -L https://github.com/${REPOSITORY_OWNER}/${APPLLICATION_REPOSITORY}/archive/${APPLICATION_VERSION}.tar.gz | tar -xz && \
    cd ${APPLLICATION_REPOSITORY}* && \
    npm install && \
    npm run build --prod


FROM maven:3.6.3-ibmjava-8-alpine AS appWebServerBuild
WORKDIR /usr/src/server
COPY . .
COPY --from=frontBuild /usr/src/${APPLLICATION_REPOSITORY}*/dist/servers ./src/main/resources/static
RUN mvn package -DskipTests


FROM openjdk:8-jdk-alpine AS appWebServerRuntime
WORKDIR /app
COPY --from=appWebServerBuild /usr/src/server/target/servers.jar ./
EXPOSE 8080
CMD java -jar servers.jar


FROM maven:3.6.3-ibmjava-8-alpine AS appServerBuild
WORKDIR /usr/src/server
COPY . .
RUN mvn package -DskipTests


FROM openjdk:8-jdk-alpine AS appServerRuntime
WORKDIR /app
COPY --from=appServerBuild /usr/src/server/target/servers.jar ./
EXPOSE 8080
CMD java -jar servers.jar
