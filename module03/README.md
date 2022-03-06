### 1.构建本地镜像

```
docker build -t httpserver:1.0 -f httpserver-dockerfile .
```

![docker build](image/01build)

![docker build](image/02build)

### 2.编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化

httpserver-dockerfile
```
FROM golang:1.17 AS build
WORKDIR /httpserver/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN go build -installsuffix cgo -o httpserver main.go

FROM busybox
COPY --from=build /httpserver/httpserver /httpserver/httpserver 
EXPOSE 8080
ENV ENV local
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]
```

### 3.将镜像推送至 docker 官方镜像仓库

```
docker login --username=15022499066 registry.cn-zhangjiakou.aliyuncs.com
docker images
docker tag 05a4486de1aa registry.cn-zhangjiakou.aliyuncs.com/sanmalove/geektime:v1.0
docker push registry.cn-zhangjiakou.aliyuncs.com/sanmalove/geektime:v1.0
```

![docker login](image/03login)
![docker push](image/05push)

### 4.通过 docker 命令本地启动 httpserver

```
docker run -d httpserver:1.0
docker ps
```

![docker run](image/06run)

### 5.通过 nsenter 进入容器查看 IP 配置

```shell
PID=$(docker inspect --format "{{ .State.Pid }}" amazing_satoshi)
nsenter -t $PID -n ip a
```

![docker nsenter](image/07nsenter)
