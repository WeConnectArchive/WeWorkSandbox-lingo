//+build mage

package main

import (
	"os"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	codePaths = []string{
		"./cmd/...",
		"./internal/...",
		"./pkg/...",
	}
)

// Run dependency downloads
func All() {
	mg.SerialDeps(Deps.InstallTools, Deps.ModDownload, Generate, Revive, Build, Test.All, Tidy)
}

type Deps mg.Namespace

// Installs pegomock and other cli tools. Mostly used for `go generate`.
func (Deps) InstallTools() error {
	var tools = []string{
		"github.com/petergtz/pegomock/pegomock",
		"github.com/mgechev/revive",
	}
	if err := runCmd("go", "install")(tools); err != nil {
		return err
	}
	return nil
}

// Runs Go Mod Download with optional debug logging
func (Deps) ModDownload() error {
	if err := run("go", "mod", "download",
		debug("-x"),
	); err != nil {
		return err
	}
	return nil
}

// Runs `go generate` with optional debug logging
func Generate() error {
	return runCmd("go", "generate",
		debug("-v"),
	)(codePaths)
}

type Test mg.Namespace

// Runs both `test:unit` and `test:acceptance` in parallel
func (Test) All() error {
	mg.Deps(Test.Unit, Test.Acceptance)
	return nil
}

// Runs all unit tests (using `-short` flag) with code coverage and optional debug logging
func (Test) Unit() error {
	pathsPlusTestArgs := append(codePaths,
		"-ginkgo.randomizeAllSpecs",
		debug("-ginkgo.progress"),
	)
	return runCmd("go", "test",
		"-short",
		"-coverprofile=unit-coverage.out",
		debug("-v"),
	)(pathsPlusTestArgs)
}

// Runs all acceptance tests with code coverage and optional debug logging
func (Test) Acceptance() error {
	pathsPlusTestArgs := append(codePaths,
		"-ginkgo.randomizeAllSpecs",
		debug("-ginkgo.progress"),
	)
	return runCmd("go", "test",
		"-coverprofile=acceptance-coverage.out",
		debug("-v"),
	)(pathsPlusTestArgs)
}

// Builds lingo and then builds codePaths
func Build() error {
	if err := run("go", "build",
		isCGOEnabled("-race"),
		debug("-v"),
		"./cmd/lingo/lingo.go",
	); err != nil {
		return err
	}
	if err := runCmd("go", "build",
		isCGOEnabled("-race"),
		debug("-v"),
	)(codePaths); err != nil {
		return err
	}
	return nil
}

// Runs `go mod tidy` with optional debug logging
func Tidy() error {
	if err := run("go", "mod", "tidy",
		debug("-v"),
	); err != nil {
		return err
	}
	return nil
}

// Run `revive` with the appropriate configs
func Revive() error {
	// NOTE: Any changes here need to be reflected in `./.github/workflows/go-revive.yml`
	if err := runCmd("revive",
		"-config", "./revive.toml",
		"-exclude", "./db/...",
		"-formatter", "stylish",
	)(codePaths); err != nil {
		return err
	}
	return nil
}

// Used to run Lingo commands
type Run mg.Namespace

// Runs Lingo Generate with default values (or config.yml if exists) and optional debug logging
func (Run) Generate() error {
	mg.SerialDeps(Build)

	return run("./lingo", "generate", debug("-v"))
}

// debug will return debugStr if mage debugging is turned on, else an empty string. Useful for enabling verbose
// output from commands.
func debug(debugStr string) string {
	return isEnabled(mg.Debug(), debugStr)
}

func isCGOEnabled(isEnabledVal string) string {
	val, ok := os.LookupEnv("CGO_ENABLED")
	return isEnabled(ok && val == "1", isEnabledVal)
}

// isEnabled will output str if b is true, else an empty string
func isEnabled(b bool, str string) string {
	if !b {
		return ""
	}
	return str
}

// run will take a normal sh.run command argument, filtering any args entries that are empty.
func run(cmd string, args ...string) error {
	args = filterEmptyStrings(args)
	return sh.Run(cmd, args...)
}

// runCmd will take a normal sh.Run command argument, and curry it with another forEachSlice argument appended,
// filtering any args or forEachSlice entries that are empty.
func runCmd(cmd string, args ...string) func(forEachSlice ...[]string) error {
	args = filterEmptyStrings(args)
	return func(forEachSlice ...[]string) error {
		for _, argSlice := range forEachSlice {
			if err := run(cmd, append(args, argSlice...)...); err != nil {
				return err
			}
		}
		return nil
	}
}

// filterEmptyStrings removes any empty strings in place
func filterEmptyStrings(strs []string) []string {
	return filterStrings(strs, func(s string) bool {
		return strings.TrimSpace(s) == ""
	})
}

// filterStrings removes any strings where filter returns true in place.
func filterStrings(strs []string, filter func(string) bool) []string {
	n := 0
	for _, x := range strs {
		if !filter(x) {
			strs[n] = x
			n++
		}
	}
	return strs[:n]
}
