package docker

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type docker struct {
	client *client.Client
}

type DockerAPI interface {
	CreateContainer() error
}

func NewDocker(dockerURL string) (docker, error) {

	apiClient, err := client.NewClientWithOpts(client.WithHost(dockerURL))
	if err != nil {
		return docker{}, err
	}

	docker := docker{client: apiClient}

	return docker, nil
}

func (docker *docker) CreateContainer() (string, error) {

	ctx := context.Background()

	config := container.Config{
		Image:        "nginx",
		Tty:          true,
		OpenStdin:    true,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
	}

	reader, err := docker.client.ImagePull(ctx, "nginx", types.ImagePullOptions{})
	if err != nil {
		return "", err
	}

	defer reader.Close()
	// cli.ImagePull is asynchronous.
	// The reader needs to be read completely for the pull operation to complete.
	// If stdout is not required, consider using io.Discard instead of os.Stdout.
	io.Copy(os.Stdout, reader)

	resp, derr := docker.client.ContainerCreate(ctx, &config, nil, nil, nil, "")
	if derr != nil {
		return "", derr
	}

	if err := docker.client.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return "", err
	}

	return resp.ID, nil
}
