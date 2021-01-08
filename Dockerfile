FROM golang

LABEL maintainer=julianlee107@hotmail.com

RUN mkdir -p /www/webapp

WORKDIR /www/webapp

COPY . /www/webapp


RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build 

EXPOSE 8080

CMD [ "./blogWithGin" ]