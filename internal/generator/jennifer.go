package generator

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dave/jennifer/jen"
)

func Render(jenFile *jen.File) (string, error) {
	buf := &bytes.Buffer{}
	if err := jenFile.Render(buf); err != nil {
		return "", transformErr(err)
	}
	return buf.String(), nil
}

// transformErr does nothing if the error is not related to formatting.
// If it is, it finds the line number,
func transformErr(e error) error {
	errStr := e.Error()

	regExp := regexp.MustCompile(`^Error (\d+).+while formatting source:`)
	subMatch := regExp.FindStringSubmatch(errStr)

	if len(subMatch) < 2 {
		return e
	}

	// Index 0 is always the full expression, so we want the first match
	lineNum, err := strconv.Atoi(subMatch[1])
	if err != nil {
		return e
	}

	var originalErrorStr = errStr[:strings.IndexRune(errStr, '\n')]
	var newErrStr = strings.Builder{}
	scanner := bufio.NewScanner(strings.NewReader(errStr))
	const linesToShow = 5
	for lineCount := 0; scanner.Scan(); lineCount++ {

		// If we are within 5 lines (before or after) the error, we want to capture that context and print that out
		if lineCount-linesToShow <= lineNum && lineCount+linesToShow >= lineNum {
			_, _ = newErrStr.WriteString(fmt.Sprintf("[Line %2d] %s\n", lineCount, scanner.Text()))
		}
	}

	return errors.New(originalErrorStr + "\n" + newErrStr.String())
}

func createBuildTag(tags ...string) string {
	//revive:disable:unhandled-error This is a string builder, never errors.
	var s strings.Builder
	s.WriteString(buildTag)
	s.WriteRune(' ')

	for idx, t := range tags {
		if idx > 0 {
			s.WriteRune(',')
		}
		s.WriteString(t)
	}
	//revive:enable:unhandled-error
	return s.String()
}
