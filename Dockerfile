FROM centos:7
COPY golang-gin /
COPY config.toml /
EXPOSE 8000
RUN chmod +x golang-gin
CMD ["/bin/bash","-l","-c","./golang-gin"]
