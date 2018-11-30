package integration_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
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
			Expect(f.Close()).To(Succeed())
		})

		It("exits zero", func() {
			Eventually(session).Should(gexec.Exit(0))
		})
	})

})
