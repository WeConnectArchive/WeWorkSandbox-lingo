package runner

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/petergtz/pegomock"
	"testing"
)

func SetupAndRun(t *testing.T, testName string) {
	// Register the global Unit Testing framework's fail handler.
	gomega.RegisterFailHandler(ginkgo.Fail)
	// Register the global Mocking framework's fail handler.
	pegomock.RegisterMockFailHandler(ginkgo.Fail)
	// Register Ginkgo
	ginkgo.RunSpecs(t, testName+" Suite")
}
