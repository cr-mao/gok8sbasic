
## go k8s 基础





```shell script
   docker pull golang:1.15.4-alpine3.12
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
 #alpine linux 中编译出 mac 的二进制文件
 # docker run --rm -it -v /Users/mac/code/gok8sbasic:/app -w /app  -e GOPROXY=https://goproxy.cn  -e CGO_ENABLED=0 -e GOOS=darwin -e GOARCH=amd64 -e GO111MODULE=on golang:1.15.4-alpine3.12 go build -o myserver  main.go
docker run --rm -it -v /Users/mac/code/gok8sbasic:/app -w /app  -e GOPROXY=https://goproxy.cn  -e CGO_ENABLED=0 -e GO111MODULE=on golang:1.15.4-alpine3.12 go build -o myserver  main.go
```
