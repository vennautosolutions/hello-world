FROM ubuntu:latest
COPY . .
EXPOSE 8081
ENTRYPOINT [ "./main" ]
