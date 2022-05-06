package sshtool

//自动连接到 服务器 去执行服务器的脚本
import (
	"bytes"
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func SSHConnect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: auth,
		// Timeout:             30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}
func Run() {
	//服务器上的脚本
	build_script := "/bin/bash /root/test.sh"
	var stdOut, stdErr bytes.Buffer
	session, err := SSHConnect("root", "xxxxx", "129.226.175.215", 22)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	session.Stdout = &stdOut
	session.Stderr = &stdErr
	err = session.Run(build_script) //也可以执行个docker 命令 如 docker run --rm -it -v /Users/mac/code/gok8sbasic:/app -w /app  -e GOPROXY=https://goproxy.cn  -e CGO_ENABLED=0 -e GO111MODULE=on golang:1.15.4-alpine3.12 go build -o myserver  main.go
	log.Println(stdOut.String())
	log.Println(stdErr.String())
	if err != nil {
		log.Fatal(err)
	}
}
