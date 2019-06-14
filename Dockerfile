FROM centos:latest
RUN mkdir /opt/app
COPY gorestapi /opt/app/gorestapi
USER root
RUN chmod +x /opt/app/gorestapi
#RUN yum -y install openssl ; yum clean all
USER 1000001
ENTRYPOINT "/opt/app/gorestapi"
