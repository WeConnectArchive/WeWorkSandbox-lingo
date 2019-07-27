package sort_test

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core Dialect
func TestSort(t *testing.T) {
	runner.SetupAndRun(t, "sort")
}
