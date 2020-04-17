package acceptance

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

func TestAcceptance(t *testing.T) {
	runner.SetupAndRun(t, "acceptance")
}
