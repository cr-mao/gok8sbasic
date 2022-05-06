package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"log"
)

//docker api 操作

/**
看这里
https://docs.docker.com/engine/api/sdk/
go get github.com/docker/docker/client

开发时一定要多看文档：
https://docs.docker.com/engine/api/latest/

*/

type Empty struct{}

func main() {
	cli, err := client.NewClient("tcp://192.168.8.110:2345", "v1.40", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	config := &container.Config{
		Cmd: []string{"./myserver"},
		ExposedPorts: map[nat.Port]struct{}{
			"8080/tcp": Empty{}, //暴露容器的8080端口
		},
		Image:      "alpine:3.12",
		WorkingDir: "/app",
	}
	hostConfig := &container.HostConfig{
		Binds: []string{"/root:/app"}, // List of volume bindings for this container
		PortBindings: map[nat.Port][]nat.PortBinding{
			"8080/tcp": []nat.PortBinding{
				nat.PortBinding{HostPort: "80"}, //宿主机的端口
			},
		},
	}
	res, err := cli.ContainerCreate(ctx, config, hostConfig, nil, nil, "myweb")
	if err != nil {
		log.Fatal(err)
	}
	err = cli.ContainerStart(ctx, res.ID, types.ContainerStartOptions{
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("容器启动成功,ID是:", res.ID)

	//images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, image := range images {
	//	fmt.Println(image.ID, image.Labels)
	//}
}
