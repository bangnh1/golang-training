package docker

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerClient interface {
	ListAll() error
	StartContainer(string) error
	StopContainer(string, *time.Duration) error
	GetContainerID(string) string
}

type CmdDocker struct {
	Ctx    context.Context
	Client *client.Client
}

func (c *CmdDocker) ListAll() error {
	conOps := types.ContainerListOptions{}
	conList, err := c.Client.ContainerList(c.Ctx, conOps)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, container := range conList {
		fmt.Printf("Container ID: %s \n", container.ID)
		fmt.Printf("Container Name: %s \n", strings.Join(container.Names, ", "))
		fmt.Printf("Image Name: %s \n", container.Image)
		fmt.Printf("Port: %s \n", portArraytoString(container.Ports))
		fmt.Printf("Status: %s \n", container.Status)
		fmt.Printf("--------------------------------- \n")
	}

	return nil
}

func (c *CmdDocker) StartContainer(containerID string) error {
	if err := c.Client.ContainerStart(c.Ctx, containerID, types.ContainerStartOptions{}); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c *CmdDocker) StopContainer(containerID string, timeout *time.Duration) error {
	if err := c.Client.ContainerStop(c.Ctx, containerID, timeout); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c *CmdDocker) GetContainerID(containerName string) string {
	conOps := types.ContainerListOptions{}
	conList, err := c.Client.ContainerList(c.Ctx, conOps)
	if err != nil {
		panic(err)
	}
	for _, container := range conList {

		IsContains := strings.Split(strings.Join(container.Names, ", "), containerName)
		if len(IsContains) > 1 {
			return container.ID
		}
	}
	return ""
}

func portArraytoString(conPorts []types.Port) string {
	var ports []string
	for _, port := range conPorts {
		ports = append(ports, port.IP+":"+strconv.Itoa(int(port.PublicPort))+":"+strconv.Itoa(int(port.PrivatePort))+":"+port.Type)
	}
	return strings.Join(ports, ", ")
}
