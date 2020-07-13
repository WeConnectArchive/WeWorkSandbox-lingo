package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/weworksandbox/lingo/internal/generator"
)

func main() {
	var dir string
	flag.StringVar(&dir, "dir", "", "absolute path to directory to generate into")
	flag.Parse()

	if err := generate(dir); err != nil {
		log.Fatal(err.Error())
	}
}

func generate(dir string) error {
	if err := validateDir(dir); err != nil {
		return err
	}
	if err := filepath.Walk(dir, generator.RemoveOldFiles(generator.GenPathFileHeader)); err != nil {
		return err
	}
	return genPaths(dir)
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

func genPaths(dir string) error {
	for _, p := range pathData {
		r, err := p.Generate()
		if err != nil {
			return fmt.Errorf("unable to generate path data '%s': %w", p.Name, err)
		}

		outputPath := filepath.Join(dir, p.Filename)
		if writeErr := generator.WriteFile(outputPath, r, os.ModePerm); writeErr != nil {
			return fmt.Errorf("unable to write to output file '%s': %w", p.Filename, writeErr)
		}
	}
	return nil
}
