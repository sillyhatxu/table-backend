FROM xushikuan/alpine-build:1.0

ENV GOPATH=/go
ENV TIME_ZONE=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

WORKDIR $GOPATH
COPY . $GOPATH
RUN mkdir -p logs
ENTRYPOINT ./main -c config.conf