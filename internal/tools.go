// +build tools

package internal

import (
	_ "github.com/petergtz/pegomock/pegomock"
)

//go:generate go install github.com/petergtz/pegomock/pegomock
