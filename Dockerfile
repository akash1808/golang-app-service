FROM centos:7
COPY golang-app-service /
COPY config.toml /
EXPOSE 8080
RUN chmod +x golang-app-service
CMD ["/bin/bash","-l","-c","./golang-app-service"]
