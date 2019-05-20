FROM centos:latest
RUN mkdir /opt/app
COPY main /opt/app/main
USER root
RUN chmod +x /opt/app/main
USER 1000001
ENTRYPOINT "/opt/app/main"
