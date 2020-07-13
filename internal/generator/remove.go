package generator

import (
	"bytes"
	"fmt"
	"os"
)

// RemoveOldFiles returns a filepath.WalkFunc to be used with filepath.Walk that removes files
// that have the first line with the same prefix as generatedPrefix.
func RemoveOldFiles(generatedPrefix string) func(string, os.FileInfo, error) error {
	pLen := len(generatedPrefix)
	data := make([]byte, pLen)

	return func(path string, info os.FileInfo, err error) (result error) {
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

		n, err := f.Read(data)
		if err != nil {
			return fmt.Errorf("unable to read file '%s' to determine if it is generated: %w", path, err)
		}

		if n != pLen { // This doesnt have enough data for us. Cannot be our file! (unless we messed up during write)
			return nil
		}
		if !bytes.EqualFold(data, []byte(generatedPrefix)) {
			return nil
		}

		if err = os.Remove(path); err != nil {
			return fmt.Errorf("unable to remove generated lingo file '%s': %w", path, err)
		}
		return nil
	}
}
