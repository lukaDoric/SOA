FROM golang:1.20 as build
WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=id=mod,type=cache,target=/go/pkg/mod \
    go mod download && go mod verify

ARG SRC=/cmd/dojo
COPY $SRC/*.go ./
RUN --mount=id=mod,type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -v -o /app/bin/app .

FROM gcr.io/distroless/base-debian11 as release-debian
ARG PORT=8080
EXPOSE $PORT
COPY --from=build /app/bin/app /app
USER nonroot:nonroot
ENTRYPOINT ["/app"]

FROM golang:1.20-alpine AS build-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=id=mod-alpine,type=cache,target=/go/pkg/mod \
    go mod download && go mod verify

ARG SRC=/cmd/dojo
COPY $SRC/*.go ./
RUN --mount=id=mod-alpine,type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -v -o /app/bin/app .

FROM alpine:3.17.3 as release-alpine
ARG PORT=8080
EXPOSE $PORT
COPY --from=build-alpine /app/bin/app /app
ENTRYPOINT ["/app"]
