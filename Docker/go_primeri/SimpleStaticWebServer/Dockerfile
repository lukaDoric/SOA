FROM alpine
WORKDIR app
COPY ./static-server .
COPY ./static ./static
RUN chmod +x ./static-server
ENTRYPOINT ["./static-server"]
