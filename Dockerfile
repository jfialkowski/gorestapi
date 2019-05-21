FROM centos:latest
RUN mkdir /opt/app
COPY gorestapi /opt/app/gorestapi
USER root
RUN chmod +x /opt/app/gorestapi
USER 1000001
ENTRYPOINT "/opt/app/gorestapi"
