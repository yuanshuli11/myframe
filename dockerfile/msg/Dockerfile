FROM golang:1.15

ARG module_name=msg
#作者
MAINTAINER Neoslyuan "1923244321@qq.com"
#设置工作目录
WORKDIR $GOPATH/src/myframe
#
#RUN  apt-get update 
#RUN  apt-get install -y supervisor
#COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
#CMD ["/usr/bin/supervisord"]
#
#RUN apt-get clean && rm -rf /var/lib/apt/lists/*
#

#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/myframe
VOLUME /work/log/myframe
#go构建可执行文件
RUN bash build.sh make $module_name
#暴露端口
EXPOSE 6064
#最终运行docker的命令
ENTRYPOINT  [ "./bin/msg", "-p","6064"]
