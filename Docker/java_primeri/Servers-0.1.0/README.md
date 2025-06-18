# Servers

Application that list servers. Data are fixed and application does not use database as a storage.

Dockerfile contains two runtime stages: 
- **appServerRuntime**  - Spring Boot application server that provides REST API.
- **appWebServerRuntime** - Spring Boot application that provides REST API and contains Angular front-end [web application](https://github.com/DanijelRadakovic/Servers-Front)

[BuildKit](https://github.com/moby/buildkit) is used for building container images.

In order to build **appServerRuntime** image run the following command (Dockerfile has to be in current working direcotry):

```shell
DOCKER_BUILDKIT=1 docker build --target appServerRuntime -t danijelradakovic/servers:0.1.0 .
```

In order to build **appServerRuntime** image run the following command (Dockerfile has to be in current working direcotry):
```shell
DOCKER_BUILDKIT=1 docker build --target appServerRuntime -t danijelradakovic/servers:0.1.0-web .
```