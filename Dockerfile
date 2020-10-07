FROM centos:7
COPY golang-gin /
COPY config.toml /
EXPOSE 8000
CMD ["./","golang-gin"]
