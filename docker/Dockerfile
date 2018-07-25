# This is my first Dockerfile
# Version 1.0
# Author: liuh

#Base images
FROM centos

#

MAINTAINER Liu H

#ADD auto tar zxf *.tar

ADD pcre-8.38.tar.gz /usr/local/src
ADD nginx-1.15.1.tar.gz /usr/local/src

#RUN

RUN yum install -y wget gcc gcc-c++ make openssl-devel
RUN useradd -s /sbin/nologin -M www

#WORKDIR
WORKDIR /usr/local/src/nginx-1.15.1

RUN ./configure --prefix=/usr/local/nginx --user=www --group=www --with-http_ssl_module --with-http_stub_status_module --with-pcre=/usr/local/src/pcre-8.38 && make && make install

RUN echo "daemon off;" >> /usr/local/nginx/conf/nginx.conf

ENV PATH /usr/local/nginx/sbin:$PATH
EXPOSE 80

CMD ["nginx"]



