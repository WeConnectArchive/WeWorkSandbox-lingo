package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/weworksandbox/lingo/internal/generate"
)

func main() {
	var dir string
	flag.StringVar(&dir, "dir", "", "absolute path to directory to generate into")
	flag.Parse()

	if err := generatePaths(dir); err != nil {
		log.Fatal(err.Error())
	}
}

func generatePaths(dir string) error {
	if err := validateDir(dir); err != nil {
		return err
	}
	if err := filepath.Walk(dir, generate.RemoveOldFiles(generate.GenPathFileHeader)); err != nil {
		return err
	}
	return genAndWrite(dir)
}

func validateDir(dir string) error {
	if !filepath.IsAbs(dir) {
		return fmt.Errorf("must be an absolute path, got: %s", dir)
	}

	if dirErr := os.MkdirAll(dir, os.ModePerm); dirErr != nil && !os.IsExist(dirErr) {
		return fmt.Errorf("unable to create directory '%s': %w", dir, dirErr)
	}

	fi, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("unable to stat directory '%s': %w", dir, err)
	}
	if !fi.IsDir() {
		return fmt.Errorf("path '%s' is not a directory", dir)
	}
	return nil
}

func genAndWrite(dir string) error {
	r, err := pathData.Generate()
	if err != nil {
		return fmt.Errorf("unable to generate path data: %w", err)
	}

	outFile := "types.go"
	outputPath := filepath.Join(dir, outFile)
	if writeErr := generate.WriteFile(outputPath, r, os.ModePerm); writeErr != nil {
		return fmt.Errorf("unable to write to output file '%s': %w", outFile, writeErr)
	}
	return nil
}
