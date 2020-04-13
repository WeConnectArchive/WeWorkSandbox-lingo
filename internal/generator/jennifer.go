package generator

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	. "github.com/dave/jennifer/jen"
	"errors"
)

func Render(jenFile *File) (string, error) {
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
	for lineCount := 0; scanner.Scan(); lineCount++ {

		// If we are within 5 lines (before or after) the error, we want to capture that context and print that out
		if lineCount-5 <= lineNum && lineCount+5 >= lineNum {
			newErrStr.WriteString(fmt.Sprintf("[Line %2d] %s\n", lineCount, scanner.Text()))
		}
	}

	return errors.New(originalErrorStr + "\n" + newErrStr.String())
}
