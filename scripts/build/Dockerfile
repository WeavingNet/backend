FROM alpine:latest
MAINTAINER zhufuyi "g.zhufuyi@gmail.com"

# set the time zone to Shanghai
RUN apk add tzdata  \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

# add curl, used for http service checking, can be installed without it if deployed in k8s
RUN apk add curl

COPY configs/ /app/configs/
COPY weaving_net /app/weaving_net
RUN chmod +x /app/weaving_net

# http port
EXPOSE 8080


WORKDIR /app

CMD ["./weaving_net", "-c", "configs/weaving_net.yml"]
# if you use the Configuration Center, weaving_net.yml is changed to the Configuration Center configuration.
#CMD ["./weaving_net", "-c", "configs/weaving_net.yml", "-enable-cc"]
