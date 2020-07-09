//+build mage

package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	// Paths and Files
	testGenerateSakilaDir            = "./internal/test/testdata/sakila"
	testGenerateSakilaConfigFileName = "lingo-config.yml"
	dockerComposeYml                 = "docker-compose.yml"

	// Go stuffs
	envCGO_ENABLED  = "CGO_ENABLED"
	cliArgSeparator = "--"

	// messages
	msgSkippingDockerCommandInCI = "Skipping - Unable to run Docker commands within CI"
)

var (
	testSchemaSakilaLingoConfigFile = filepath.Join(testGenerateSakilaDir, testGenerateSakilaConfigFileName)
	testGenerateSakilaDC            = filepath.Join(testGenerateSakilaDir, dockerComposeYml)
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
		Gen.TestSchema,
		// Keep GoVet, Build & Test after Gen.TestSchema. If its before, the generated schema files are included
		// in the building / vet / tests. If there are contract changes (due to development), it will fail to compile.
		GoVet,
		Build,
		Test.All,
		GoTidy,
		Gen.StopTestSchemaDB,
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
		debug("-x"),
	)(allFiles); err != nil {
		return err
	}
	return nil
}

type Test mg.Namespace

// Runs both `test:unit`, `test:functional` and `test:benchmark` in parallel. If not in CI, it will start the
// test schema DB, generate the schema, run the tests, and stop the db.
func (Test) All() error {
	mg.SerialDeps(Test.Unit, Test.Functional, Test.Benchmark, Test.IntegrationWithDB)
	return nil
}

// Runs all unit tests (using `-short` flag) with code coverage and optional debug logging
func (Test) Unit() error {
	pathsPlusTestArgs := append(codePaths,
		cliArgSeparator,
		"--ginkgo.randomizeAllSpecs",
		debug("--ginkgo.progress"),
	)
	return runCmd("go", "test",
		"-short",
		"-coverprofile=unit-coverage.out",
		goRaceFlag(),
		debug("-v"),
	)(pathsPlusTestArgs)
}

// Runs all functional tests with code coverage and optional debug logging
func (Test) Functional() error {
	pathsPlusTestArgs := append(codePaths,
		cliArgSeparator,
		"--ginkgo.randomizeAllSpecs",
		debug("--ginkgo.progress"),
	)
	return runCmd("go", "test",
		"-short",
		"-coverprofile=functional-coverage.out",
		goRaceFlag(),
		debug("-v"),
	)(pathsPlusTestArgs)
}

// Runs just the integration tests against an already running DB instance.
func (Test) Integration() error {
	absConfig, err := filepath.Abs(testSchemaSakilaLingoConfigFile)
	if err != nil {
		return fmt.Errorf("unable to find absolute path for config file: %w", err)
	}

	pathsPlusTestArgs := append(codePaths,
		cliArgSeparator,
		"--ginkgo.randomizeAllSpecs",
		debug("--ginkgo.progress"),
		"--",
		"--config", absConfig,
	)
	return runCmd("go", "test",
		"-coverprofile=integration-coverage.out",
		goRaceFlag(),
		debug("-v"),
	)(pathsPlusTestArgs)
}

// Runs the test schema database, the integration tests, and then cleans up the DB. Does nothing if in CI.
func (Test) IntegrationWithDB() {
	if isCI() {
		log.Println("Cannot run docker containers in CI")
		return
	}

	mg.SerialDeps(Gen.StartTestSchemaDB, Test.Integration)
}

// Runs all benchmark tests with -benchmem
func (Test) Benchmark() error {
	if err := runCmd("go", "test",
		"-run=XXX", // Random fake test name
		"-bench", "Benchmark",
		"-short",
		"-benchmem",
		debug("-v"),
	)(allPkgs); err != nil {
		return err
	}
	return nil
}

// Builds lingo CLI and then builds codePaths
func Build() error {
	if err := run("go", "build",
		goRaceFlag(),
		debug("-v"),
		"./cmd/lingo/lingo.go",
	); err != nil {
		return err
	}
	if err := runCmd("go", "build",
		goRaceFlag(),
		debug("-v"),
	)(codePaths); err != nil {
		return err
	}
	return nil
}

// Builds lingo CLI and then builds all codePaths without building any generated lingo files.
func BuildNoLingoGens() error {
	if err := run("go", "build",
		goRaceFlag(),
		debug("-v"),
		"./cmd/lingo/lingo.go",
	); err != nil {
		return err
	}
	if err := runCmd("go", "build",
		goRaceFlag(),
		debug("-v"),
		"-tags",
		"nolingo",
	)(codePaths); err != nil {
		return err
	}
	return nil
}

// Starts the test db, installs the schema, then runs lingo to generate the files
func (Gen) TestSchema() error {
	mg.SerialDeps(Gen.StartTestSchemaDB, Run.LingoGenTestSchema)
	return nil
}

// This will start the DB then installs the schema & data.
// The test schema and data comes from MySQL: https://dev.mysql.com/doc/index-other.html
func (Gen) StartTestSchemaDB(ctx context.Context) error {
	if isCI() {
		log.Println(msgSkippingDockerCommandInCI)
		return nil
	}

	started := time.Now()
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

	timeout := 30 * time.Second
	log.Printf("Waiting up to %s for schema script to complete", timeout)
	if err := waitForContainerExit(ctx, "recreate-schema-script", started, 1*time.Second, timeout, func() {
		fmt.Print(".")
	}); err != nil {
		return fmt.Errorf("failed to run test schema script successfully: %w", err)
	}
	log.Println("\nDatabase setup completed")
	return nil
}

// Stop the test containers used to sakila our schema
func (Gen) StopTestSchemaDB() error {
	if isCI() {
		log.Println(msgSkippingDockerCommandInCI)
		return nil
	}

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
func GoTidy() error {
	if err := run("go", "mod", "tidy",
		debug("-v"),
	); err != nil {
		return err
	}
	return nil
}

// Runs `go vet` with optional debug logging
func GoVet() error {
	if err := runCmd("go", "vet", debug("-v"))(allPkgs); err != nil {
		return err
	}
	return nil
}

// Runs `revive` with the appropriate configs
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
	mg.SerialDeps(BuildNoLingoGens)

	if isCI() {
		log.Println(msgSkippingDockerCommandInCI)
		return nil
	}

	return run("./lingo", "generate",
		"--config", testSchemaSakilaLingoConfigFile,
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

// isCGOEnabled returns true if CGO_ENABLED is not found in env vars, or if its value is equal to string "1".
func isCGOEnabled() bool {
	val, ok := os.LookupEnv(envCGO_ENABLED)
	return !ok || val == "1" // If it is not found, its enabled by default in Go
}

// ifCGOEnabled returns returnIfEnabled if CGO_ENABLED is true
func ifCGOEnabled(returnIfEnabled string) string {
	if !isCGOEnabled() {
		return ""
	}
	return returnIfEnabled
}

// isCI returns true if an env var for the CI system enabled is present
func isCI() bool {
	val, ok := os.LookupEnv("GITHUB_ACTIONS")
	return ok && val == "true"
}

// goRaceFlag will return -race if CGO_ENABLED is true
func goRaceFlag() string {
	const flagRace = "-race"
	return ifCGOEnabled(flagRace)
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

func waitForContainerExit(
	ctx context.Context, name string, started time.Time, skew time.Duration, timeout time.Duration, onTick func(),
) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return fmt.Errorf("unable to create docker Go SDK client from env vars: %w", err)
	}

	listArgs := filters.NewArgs()
	listArgs.Add("name", name)

	timeoutCtx, _ := context.WithTimeout(ctx, timeout)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	ensureSkewDirection := time.Duration(math.Abs(float64(skew)))
	for {
		select {
		case <-ticker.C:
			if onTick != nil {
				onTick()
			}

			clo := types.ContainerListOptions{
				Latest:  true,
				All:     true,
				Filters: listArgs,
			}

			containers, err := cli.ContainerList(context.Background(), clo)
			if err != nil {
				return fmt.Errorf("unable to query if docker container %s is completed: %w", name, err)
			}

			for _, container := range containers {
				createdAt := time.Unix(container.Created, 0)

				if createdAt.Before(started.Add(-ensureSkewDirection)) {
					continue
				}
				if !strings.EqualFold(container.State, "exited") {
					continue
				}
				if !strings.HasPrefix(container.Status, "Exited (0)") {
					return fmt.Errorf("expected contianer to exit with status prefix 'Exited (0)', "+
						"but container exited with status %s", container.Status)
				}
				return nil
			}
		case <-timeoutCtx.Done():
			return timeoutCtx.Err()
		}
	}
}
