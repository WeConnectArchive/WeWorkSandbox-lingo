package runner

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/petergtz/pegomock"
)

func SetupAndRunUnit(t *testing.T, unitTestName, testCategory string) {
	// If we aren't running unit, why run unit?
	if !testing.Short() {
		return
	}

	// Register the global Unit Testing framework's fail handler.
	gomega.RegisterFailHandler(ginkgo.Fail)
	// Register the global Mocking framework's fail handler.
	pegomock.RegisterMockFailHandler(ginkgo.Fail)
	// Register Ginkgo
	ginkgo.RunSpecs(t, unitTestName+" Suite - "+testCategory)
}

func SetupAndRunFunctional(t *testing.T, funcTestName string) {
	// These are long-er running.
	if testing.Short() {
		return
	}

	// Register the global test framework's fail handler.
	gomega.RegisterFailHandler(ginkgo.Fail)
	// Register Ginkgo
	ginkgo.RunSpecs(t, funcTestName+" Suite - Functional")
}
