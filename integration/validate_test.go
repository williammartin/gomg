package integration_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/williammartin/omg"
)

var _ = Describe("gomg validate", func() {

	var (
		gomgCmd *exec.Cmd

		session *gexec.Session
	)

	BeforeEach(func() {
		gomgCmd = exec.Command(gomgPath)
		gomgCmd.Args = append(gomgCmd.Args, "validate")
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
			BeforeEach(func() {
				microservice := generateValidMicroservice()
				y, err := yaml.Marshal(microservice)
				Expect(err).NotTo(HaveOccurred())
				Expect(ioutil.WriteFile(filepath.Join(tmpDir, "microservice.yml"), y, 0777)).To(Succeed())
			})

			It("exits zero", func() {
				Eventually(session).Should(gexec.Exit(0))
			})

			It("displays a message that validation succeeded", func() {
				Eventually(session).Should(gbytes.Say("validation succeeded"))
			})

			It("displays a success message", func() {
				Eventually(session).Should(gbytes.Say("SUCCESS"))
			})
		})
	})
})

func generateValidMicroservice() *omg.Microservice {
	return &omg.Microservice{
		OMG: 1,
		Info: &omg.Info{
			Version:     "0.0.1",
			Title:       "Test Microservice",
			Description: "A Test Microservice",
			License: &omg.License{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
		},
	}
}
