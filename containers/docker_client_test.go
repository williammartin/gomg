package containers_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/onsi/gomega/gbytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/williammartin/gomg/containers"
)

var _ = Describe("Docker Client", func() {
	var (
		client *DockerClient
	)

	BeforeEach(func() {
		endpoint := "unix:///var/run/docker.sock"
		dockerClient, err := docker.NewClient(endpoint)
		Expect(err).NotTo(HaveOccurred())

		client = &DockerClient{
			DockerClient: dockerClient,
		}
	})

	Describe("Build", func() {
		When("provided with valid arguments", func() {
			var (
				tempDir string

				repository string
				tag        string
				name       string
			)

			BeforeEach(func() {
				var err error
				tempDir, err = ioutil.TempDir("", "")
				Expect(err).NotTo(HaveOccurred())

				dockerfile := filepath.Join(tempDir, "Dockerfile")
				ioutil.WriteFile(dockerfile, []byte("FROM scratch\nSTOPSIGNAL 9"), 0600)

				repository = prefixedRandomName("img")
				tag = prefixedRandomName("tag")
				name = fmt.Sprintf("%s:%s", repository, tag)
			})

			AfterEach(func() {
				cmd := exec.Command("docker", "rmi", name)
				cmd.Stdout = GinkgoWriter
				cmd.Stderr = GinkgoWriter
				Expect(cmd.Run()).To(Succeed())

				Expect(os.RemoveAll(tempDir)).To(Succeed())
			})

			It("builds an image with the given name", func() {
				err := client.Build(name, WithContextDir(tempDir), WithOutputStream(ioutil.Discard))
				Expect(err).NotTo(HaveOccurred())

				stdout := gbytes.NewBuffer()
				cmd := exec.Command("docker", "images", repository)
				cmd.Stdout = io.MultiWriter(GinkgoWriter, stdout)
				cmd.Stderr = GinkgoWriter
				Expect(cmd.Run()).To(Succeed())

				Expect(stdout).To(gbytes.Say(repository))
				Expect(stdout).To(gbytes.Say(tag))
			})

			It("forwards output to the provided writer", func() {
				output := gbytes.NewBuffer()
				err := client.Build(name, WithContextDir(tempDir), WithOutputStream(output))
				Expect(err).NotTo(HaveOccurred())
				Expect(output).To(gbytes.Say(`.*`))
			})
		})

		When("building the image fails", func() {
			It("returns an error", func() {
				Expect(client.Build("!invalid-name!")).NotTo(Succeed())
			})
		})
	})
})

func prefixedRandomName(namePrefix string) string {
	return namePrefix + "-" + randomName()
}

func randomName() string {
	guid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return strings.Replace(guid.String(), "-", "", -1)
}
