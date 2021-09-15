package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strconv"
	"time"
)

func Docker() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	args := os.Args
	timeout := time.Duration(60)

	switch args[1] {
	case "ps":
		listAllCons(ctx, cli)
	case "start":
		for _, container := range args[2:] {
			containerID := getContainerID(ctx, cli, container)
			if containerID == "" {
				StartCon(ctx, cli, container)
			} else {
				StartCon(ctx, cli, containerID)
			}
		}
	case "stop":
		for _, container := range args[2:] {
			containerID := getContainerID(ctx, cli, container)
			if containerID == "" {
				StopCon(ctx, cli, container, &timeout)
			} else {
				StopCon(ctx, cli, containerID, &timeout)
			}
		}
	}
}

func portArraytoString(conPorts []types.Port) string {
	var ports []string
	for _, port := range conPorts {
		ports = append(ports, port.IP+":"+strconv.Itoa(int(port.PublicPort))+":"+strconv.Itoa(int(port.PrivatePort))+":"+port.Type)
	}
	return strings.Join(ports, ", ")
}

func listAllCons(ctx context.Context, cli *client.Client) {

	conOps := types.ContainerListOptions{}
	conList, err := cli.ContainerList(ctx, conOps)
	if err != nil {
		panic(err)
	}
	// fmt.Println(conList)
	for _, container := range conList {
		fmt.Printf("Container ID: %s \n", container.ID)
		fmt.Printf("Container Name: %s \n", strings.Join(container.Names, ", "))
		fmt.Printf("Image Name: %s \n", container.Image)
		fmt.Printf("Port: %s \n", portArraytoString(container.Ports))
		fmt.Printf("Status: %s \n", container.Status)
		fmt.Printf("--------------------------------- \n")
	}
}

func StartCon(ctx context.Context, cli *client.Client, containerID string) {
	if err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
}

func StopCon(ctx context.Context, cli *client.Client, containerID string, timeout *time.Duration) {
	if err := cli.ContainerStop(ctx, containerID, timeout); err != nil {
		panic(err)
	}
}

func getContainerID(ctx context.Context, cli *client.Client, containerName string) string {
	conOps := types.ContainerListOptions{}
	conList, err := cli.ContainerList(ctx, conOps)
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
