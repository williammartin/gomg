package containers

import (
	"fmt"
	"io"
	"path/filepath"

	docker "github.com/fsouza/go-dockerclient"
)

type DockerClient struct {
	DockerClient *docker.Client
}

func (c *DockerClient) Build(name, tag, dockerfile string, writer io.Writer) error {
	opts := docker.BuildImageOptions{
		Name:         fmt.Sprintf("%s:%s", name, tag),
		OutputStream: writer,
		ContextDir:   filepath.Dir(dockerfile),
	}

	return c.DockerClient.BuildImage(opts)
}
