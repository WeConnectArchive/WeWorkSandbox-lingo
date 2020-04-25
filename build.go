//+build mage

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	// Paths and Files
	testGenerateSakilaDir = "./internal/test/testdata/sakila"
	dockerComposeYml      = "docker-compose.yml"
)

var (
	testGenerateSakilaDC = filepath.Join(testGenerateSakilaDir, dockerComposeYml)
	// allFiles has the current directory `.` plus all subsequent directories `./...`.
	// Note: `...` does not work as it tries to do EVERYTHING including random GOPATH stuff.
	allFiles = append(allPkgs, ".")
	allPkgs  = []string{
		"./...",
	}
	codePaths = []string{
		"./cmd/...",
		"./internal/...",
		"./pkg/...",
	}
)

// Run all the things that CI does
func All() {
	s := time.Now()
	log.Printf("Starting Build - %s", s.Format("15:04:05.999999999"))
	defer func() {
		if err := recover(); err != nil {
			log.Printf("ERROR - Took %s", time.Since(s))
			panic(err)
		} else {
			log.Printf("Completed - Took %s", time.Since(s))
		}
	}()

	mg.SerialDeps(
		Deps.InstallTools,
		Deps.ModDownload,
		Gen.Go,
		GoFmt,
		Revive,
		Build,
		Test.All,
		Tidy,
	)
}

type Deps mg.Namespace

// Installs pegomock and other cli tools. Mostly used for `go generate`.
func (Deps) InstallTools() error {
	var tools = []string{
		"github.com/petergtz/pegomock/pegomock",
		"github.com/mgechev/revive",
	}
	if err := runCmd("go", "install", debug("-x"))(tools); err != nil {
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

type Gen mg.Namespace

// Runs `go generate` with optional debug logging
func (Gen) Go() error {
	return runCmd("go", "generate",
		debug("-v"),
	)(codePaths)
}

// Runs `go fmt` with optional debug logging
func GoFmt() error {
	if err := runCmd("go", "fmt",
		debug("-v"),
	)(allFiles); err != nil {
		return err
	}
	return nil
}

type Test mg.Namespace

// Runs both `test:unit` and `test:acceptance` in parallel
func (Test) All() error {
	mg.Deps(Test.Unit, Test.Functional, Test.Benchmark)
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

// Runs all functional tests with code coverage and optional debug logging
func (Test) Functional() error {
	pathsPlusTestArgs := append(codePaths,
		"-ginkgo.randomizeAllSpecs",
		debug("-ginkgo.progress"),
	)
	return runCmd("go", "test",
		"-short",
		"-coverprofile=functional-coverage.out",
		debug("-v"),
	)(pathsPlusTestArgs)
}

// Runs all benchmark tests with -benchmem
func (Test) Benchmark() error {
	if err := runCmd("go", "test",
		"-bench", "Benchmark.",
		"-benchmem",
	)(allPkgs); err != nil {
		return err
	}
	return nil
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

// Starts the test db, installs the schema, then runs lingo to generate the files. If successful, the db
// is shutdown and deleted.
func (Gen) TestSchema() error {
	mg.SerialDeps(Gen.StartTestSchemaDB, Run.LingoGenTestSchema, Gen.StopTestSchemaDB)
	return nil
}

// This will start the DB then installs the schema & data.
// The test schema and data comes from MySQL: https://dev.mysql.com/doc/index-other.html
func (Gen) StartTestSchemaDB() error {
	if isCI() {
		log.Println("Skipping - Unable to run Docker commands within CI")
		return nil
	}
	if err := run("docker-compose",
		"-f", testGenerateSakilaDC,
		"--env-file", filepath.Join(testGenerateSakilaDir, ".env"),
		"up",
		"-d",
		"--remove-orphans",
		"--force-recreate",
	); err != nil {
		return err
	}
	time.Sleep(3 * time.Second)
	log.Println("Database should be completed")
	return nil
}

// Stop the test containers used to sakila our schema
func (Gen) StopTestSchemaDB() error {
	if err := run("docker-compose",
		"-f", testGenerateSakilaDC,
		"rm",
		"--stop",
		"--force",
		"-v",
	); err != nil {
		_ = mg.Fatalf(1, "Unable to remove docker container test schema database")
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
	// Allows for Cmd + Click on the line in logs to go directly to it
	var formatter = "friendly"
	if isCI() {
		// Groups each files errors together for easy lint fixing.
		formatter = "stylish"
	}

	// NOTE: Any changes here need to be reflected in `./.github/workflows/go-revive.yml`
	if err := runCmd("revive",
		"-config", "./revive.toml",
		"-exclude", "./db/...",
		"-exclude", "./internal/test/schema/...",
		"-formatter", formatter,
	)(codePaths); err != nil {
		return err
	}
	return nil
}

// Used to run Lingo commands
type Run mg.Namespace

// Runs Lingo for the Sakila test DB
func (Run) LingoGenTestSchema() error {
	mg.SerialDeps(Build)

	return run("./lingo", "generate",
		"--config", filepath.Join(testGenerateSakilaDir, "lingo-config.yml"),
	)
}

// debug will return debugStr if mage debugging is turned on, else an empty string. Useful for enabling verbose
// output from commands.
func debug(debugStr string) string {
	if mg.Debug() {
		return debugStr
	}
	return ""
}

// isCGOEnabled returns returnIfEnabled if CGO_ENABLED is true
func isCGOEnabled(returnIfEnabled string) string {
	val, ok := os.LookupEnv("CGO_ENABLED")
	if ok && val == "1" {
		return returnIfEnabled
	}
	return ""
}

// isCI returns true if an env var for the CI system enabled is present
func isCI() bool {
	val, ok := os.LookupEnv("GITHUB_ACTIONS")
	return ok && val == "true"
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

func downloadFile(url string, tempFilePattern string) (string, int64, error) {
	c := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := c.Get(url)
	if err != nil {
		return "", 0, fmt.Errorf("unable to download test database: %w", err)
	}
	if http.StatusOK != resp.StatusCode {
		return "", 0, fmt.Errorf("invalid response code while trying to download file: %w", err)
	}
	fileName, size, err := copyToTempFile(resp.Body, tempFilePattern)
	if err != nil {
		return "", 0, fmt.Errorf("")
	}
	return fileName, size, nil
}

func copyToTempFile(r io.ReadCloser, tempFilePattern string) (string, int64, error) {
	outFile, err := ioutil.TempFile("", tempFilePattern)
	if err != nil {
		return "", 0, fmt.Errorf("unable to create temp file: %w", err)
	}
	size, err := io.Copy(outFile, r)
	if err != nil {
		return "", 0, fmt.Errorf("unable to copy data to temp file: %w", err)
	}
	if err := r.Close(); err != nil {
		return "", 0, fmt.Errorf("unable to close reader: %w", err)
	}
	if err := outFile.Close(); err != nil {
		return "", 0, fmt.Errorf("unable to close file: %w", err)
	}
	return outFile.Name(), size, nil
}
