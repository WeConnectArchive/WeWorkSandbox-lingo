package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

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
	if err := filepath.Walk(dir, removeOldFile); err != nil {
		return err
	}
	return jen(dir)
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

func removeOldFile(path string, info os.FileInfo, err error) (result error) {
	if err != nil {
		return fmt.Errorf("error while accessing path %s: %w", path, err)
	}
	if info.IsDir() {
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("unable to open file '%s' to determine if it is generated: %w", path, err)
	}
	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		if scanErr := scanner.Err(); scanErr != nil {
			return fmt.Errorf("unable to read file '%s' to determine if it is generated: %w", path, err)
		}
		return nil
	}

	if line := scanner.Text(); !strings.EqualFold(line, generator.GenPathFileHeader) {
		return nil
	}

	if err = os.Remove(path); err != nil {
		return fmt.Errorf("unable to remove old lingo path file: %w", err)
	}
	return nil
}

func jen(dir string) error {

	for _, p := range pathData {
		r, err := p.Generate()
		if err != nil {
			return fmt.Errorf("unable to generate path '%s': %w", p.Name, err)
		}

		if writeErr := writeFile(dir, p.Filename, r); writeErr != nil {
			return fmt.Errorf("unable to write path file '%s': %w", p.Filename, writeErr)
		}
	}
	return nil
}

func writeFile(dir, pathName string, data io.Reader) error {
	path := filepath.Join(dir, pathName)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to open file for write: %w", err)
	}
	if _, err = io.Copy(f, data); err != nil {
		return fmt.Errorf("unable to copy file data: %w", err)
	}
	if err = f.Close(); err != nil {
		return fmt.Errorf("unable to close file: %w", err)
	}
	return nil
}
