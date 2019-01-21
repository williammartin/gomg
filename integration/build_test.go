package integration_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/williammartin/omg"
)

var _ = Describe("gomg build", func() {
	var (
		gomgCmd *exec.Cmd

		session *gexec.Session
	)

	BeforeEach(func() {
		gomgCmd = exec.Command(gomgPath)
		gomgCmd.Args = append(gomgCmd.Args, "build")
	})

	JustBeforeEach(func() {
		session = execBin(gomgCmd)
	})

	var tmpDir string

	BeforeEach(func() {
		var err error
		tmpDir, err = ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())

		gomgCmd.Dir = tmpDir
	})

	AfterEach(func() {
		Expect(os.RemoveAll(tmpDir)).To(Succeed())
	})

	When("the current directory does not contain a microservice.yml", func() {
		It("exits non-zero", func() {
			Eventually(session).Should(gexec.Exit(1))
		})

		It("displays an informative error", func() {
			Eventually(session.Err).Should(gbytes.Say("the current directory must contain a 'microservice.yml' file"))
		})

		It("displays FAILED", func() {
			Eventually(session.Err).Should(gbytes.Say("FAILED"))
		})
	})

	When("the current directory contains a microservice.yml", func() {
		BeforeEach(func() {
			f, err := os.Create(filepath.Join(tmpDir, "microservice.yml"))
			Expect(err).NotTo(HaveOccurred())

			f.WriteString("fake: key")
			Expect(f.Close()).To(Succeed())
		})

		When("the microservice.yml is invalid (because it is empty)", func() {
			It("exits non-zero", func() {
				Eventually(session).Should(gexec.Exit(1))
			})

			It("displays useful validation errors", func() {
				Eventually(session.Err).Should(gbytes.Say("validation errors occurred:"))
				Eventually(session.Err).Should(gbytes.Say("omg is required"))
			})

			It("displays that the command failed", func() {
				Eventually(session.Err).Should(gbytes.Say("FAILED"))
			})
		})

		When("the microservice.yml is well formed", func() {
			var microservice *omg.Microservice

			BeforeEach(func() {
				microservice = generateValidMicroservice()
				y, err := yaml.Marshal(microservice)
				Expect(err).NotTo(HaveOccurred())
				Expect(ioutil.WriteFile(filepath.Join(tmpDir, "microservice.yml"), y, 0777)).To(Succeed())
			})

			When("there is a Dockerfile present", func() {
				BeforeEach(func() {
					copyFile("assets/Dockerfile", filepath.Join(tmpDir, "Dockerfile"))
				})

				AfterEach(func() {
					expectedName := fmt.Sprintf("omg-%s", strings.ToLower(strings.Replace(microservice.Info.Title, " ", "", -1)))
					cmd := exec.Command("docker", "rmi", fmt.Sprintf("%s:%s", expectedName, "latest"))
					cmd.Stdout = GinkgoWriter
					cmd.Stderr = GinkgoWriter
					Expect(cmd.Run()).To(Succeed())
				})

				It("builds a docker image with the title of the microservice and the latest tag", func() {
					expectedName := fmt.Sprintf("omg-%s", strings.ToLower(strings.Replace(microservice.Info.Title, " ", "", -1)))
					Eventually(session).Should(gexec.Exit(0))
					Eventually(session).Should(gbytes.Say(`building\.\.\.`))
					Eventually(session).Should(gbytes.Say("Successfully built"))
					Eventually(session).Should(gbytes.Say("built %s with tag latest", expectedName))
					Eventually(session).Should(gbytes.Say("SUCCESS"))

					stdout := gbytes.NewBuffer()
					cmd := exec.Command("docker", "images", fmt.Sprintf("%s:%s", expectedName, "latest"))
					cmd.Stdout = io.MultiWriter(GinkgoWriter, stdout)
					cmd.Stderr = GinkgoWriter
					Expect(cmd.Run()).To(Succeed())

					Expect(stdout).To(gbytes.Say(expectedName))
				})
			})

			When("there is no Dockerfile present", func() {
				It("displays an informative error", func() {
					Eventually(session.Err).Should(gbytes.Say("the current directory must contain a 'Dockerfile' file"))
				})
			})
		})
	})
})

func copyFile(src string, dst string) {
	data, err := ioutil.ReadFile(src)
	Expect(err).NotTo(HaveOccurred())
	Expect(ioutil.WriteFile(dst, data, 0644)).To(Succeed())
}
