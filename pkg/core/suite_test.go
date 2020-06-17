package core_test

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

func TestCore(t *testing.T) {
	runner.SetupAndRunUnit(t, "core", "unit")
}

