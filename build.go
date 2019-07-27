//+build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func All() {
	mg.Deps(Deps, Test.Unit, Test.Acceptance)
}

// Manage your deps, or running package managers.
func Deps() error {
	fmt.Println("Installing Deps...")
	return sh.Run("go", "get", "github.com/petergtz/pegomock")
}

type Gen mg.Namespace

func (Gen) Mocks() error {
	mg.Deps(Deps)
	return sh.Run("go", "generate", "-v", "./pkg/...")
}

type Test mg.Namespace

func (Test) Unit() error {
	mg.Deps(Gen.Mocks)
	goTest := RunCmd("go", "test", "-v", "-count=1")
	return goTest([]string{
		"./internal/...",
		"./pkg/...",
	})
}

func (Test) Acceptance() error {
	return nil
}

func Build() error {
	return sh.Run("go", "build", "-race", "-v", "./cmd/lingo/lingo.go")
}

type Run mg.Namespace

func (Run) Generate() error {
	mg.SerialDeps(Build)

	return sh.Run("./lingo", "generate")
}

// Clean up after yourself
func Clean() error {
	mg.SerialDeps(Deps)
	return sh.Run("pegomock", "remove", "-n", "-r")
}

func RunCmd(cmd string, args ...string) func(...[]string) error {
	return func(args2 ...[]string) error {
		for _, arg := range args2 {
			if err := sh.Run(cmd, append(args, arg...)...); err != nil {
				return err
			}
		}
		return nil
	}
}
