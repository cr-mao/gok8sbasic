## go k8s 基础




#### 编译二进制文件，用另外一个容器启动这个二进制
```shell script
   docker pull golang:1.15.4-alpine3.12
# CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
 #alpine linux 中编译出 mac 的二进制文件 ，mac 用这个 
 # docker run --rm -it -v /Users/mac/code/gok8sbasic:/app -w /app  -e GOPROXY=https://goproxy.cn  -e CGO_ENABLED=0 -e GOOS=darwin -e GOARCH=amd64 -e GO111MODULE=on golang:1.15.4-alpine3.12 go build -o myserver  main.go

#编译出二进制文件即可
docker run --rm -it -v /Users/mac/code/gok8sbasic:/app -w /app  -e GOPROXY=https://goproxy.cn  -e CGO_ENABLED=0 -e GO111MODULE=on golang:1.15.4-alpine3.12 go build -o myserver  main.go
# 
docker run --name myweb -d -v /Users/mac/code/gok8sbasic:/app -w /app -p 80:8080 alpine:3.12 ./myserver
```


#### 开发docker tcp 连接
```text 

对于centos7  文件在
/usr/lib/systemd/system/docker.service

找到这个 （注释掉）
# ExecStart=/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock 
改成：
ExecStart=/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock -H tcp://0.0.0.0:2345

systemctl restart docker.service
systemctl daemon-reload  
```

#### go脚本创建启动一个容器

