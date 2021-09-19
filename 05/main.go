package main

import (
	"context"
	"os"
	"time"

	"github.com/bangnh1/golang-training/05/docker"
	"github.com/docker/docker/client"
)

func main() {
	Docker()
	var pemp1 Employee
	pemp1 = &Permanent{1, 1000, 50}
	pemp2 := pemp1.Clone()
	pemp3 := pemp1.Clone()
	var cemp1 Employee
	cemp1 = &Contract{4, 3000}
	employees := []Employee{pemp1, pemp2, pemp3, cemp1}
	totalExpense(employees)
}

func Docker() {

	newClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	var cmd docker.DockerClient
	cmd = &docker.CmdDocker{
		Ctx:    context.Background(),
		Client: newClient,
	}

	args := os.Args
	timeout := time.Duration(60)

	switch args[1] {
	case "ps":
		cmd.ListAll()
	case "start":
		for _, container := range args[2:] {
			containerID := cmd.GetContainerID(container)
			if containerID == "" {
				cmd.StartContainer(container)
			} else {
				cmd.StartContainer(containerID)
			}
		}
	case "stop":
		for _, container := range args[2:] {
			containerID := cmd.GetContainerID(container)
			if containerID == "" {
				cmd.StopContainer(container, &timeout)
			} else {
				cmd.StopContainer(containerID, &timeout)
			}
		}
	}
}
