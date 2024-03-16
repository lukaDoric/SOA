# Servers

Application that manages servers. Data are stored in relational database. Application has 3 profiles and use different database in each profile:
- dev - use MYSQL 8.0.19
- test - use POSTGRES 13
- prod - use POSTGRES 13

Dockerfile contains two runtime stages: 
- **appServerRuntime**  - Spring Boot application server that provides REST API.
- **appWebServerRuntime** - Spring Boot application that provides REST API and contains Angular front-end [web application](https://github.com/DanijelRadakovic/Servers-Front)

[BuildKit](https://github.com/moby/buildkit) is used for building container images.

In order to build **appServerRuntime** image run the following command (Dockerfile has to be in current working direcotry and allowed values for stage are dev, test, prod):

```shell
DOCKER_BUILDKIT=1 docker build --target appServerRuntime --build-args STAGE=dev -t danijelradakovic/servers:0.2.0 .
```

In order to build **appServerRuntime** image run the following command (Dockerfile has to be in current working direcotry and allowed values for stage are dev, test, prod):
```shell
DOCKER_BUILDKIT=1 docker build --target appServerRuntime --build-args STAGE=dev -t danijelradakovic/servers:0.2.1 .
```

Building container images can also be achieved using docker compose. Before running any docker compose command you should always check configuration using the following command:
```shell
docker-compose --env-file config/.env.dev config
```

To setup an infrastructure run the following command:
```shell
docker-compose --env-file config/.env.dev up --build
```

To setup an infrastructure for test environment run the fllowing command:
```shell
docker-compose --env-file config/.env.test -f docker-compose.yml -f docker-compose.test.yml up --build
```

To setup an infrastructure for production environment run the fllowing command:
```shell
docker-compose --env-file config/.env.prod -f docker-compose.yml -f docker-compose.prod.yml up --build
```
