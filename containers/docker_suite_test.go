package containers_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestContainers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Containers Suite")
}
