FROM ubuntu:latest
COPY . .
EXPOSE 8080
ENTRYPOINT [ "./main" ]
