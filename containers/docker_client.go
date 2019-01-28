package containers

import (
	"io"

	docker "github.com/fsouza/go-dockerclient"
)

type DockerClient struct {
	DockerClient *docker.Client
}

type BuildOptions struct {
	Name         string
	ContextDir   string
	OutputStream io.Writer
}

type BuildOpt func(buildOpts *BuildOptions)

func WithContextDir(dirPath string) BuildOpt {
	return func(buildOpts *BuildOptions) {
		buildOpts.ContextDir = dirPath
	}
}

func WithOutputStream(writer io.Writer) BuildOpt {
	return func(buildOpts *BuildOptions) {
		buildOpts.OutputStream = writer
	}
}

func (c *DockerClient) Build(name string, opts ...BuildOpt) error {
	buildOpts := &BuildOptions{Name: name}
	for _, o := range opts {
		o(buildOpts)
	}

	dockerBuildOpts := convertToDockerBuildOpts(buildOpts)
	return c.DockerClient.BuildImage(dockerBuildOpts)
}

func convertToDockerBuildOpts(buildOpts *BuildOptions) docker.BuildImageOptions {
	return docker.BuildImageOptions{
		Name:         buildOpts.Name,
		ContextDir:   buildOpts.ContextDir,
		OutputStream: buildOpts.OutputStream,
	}
}
