package bootstrap

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// NewFileInstructions takes a file path and return its content into []string or an error
func NewFileInstructions(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file path %s, got %q", path, err)
	}

	defer f.Close()

	lines, err := contentToStringArray(f)
	if err != nil {
		return nil, fmt.Errorf("failed to conver file content to array of strings, got %q", err)
	}

	return lines, nil
}

// contentToStringArray takes an io.Reader and convert each lines/row/bytes into string returning a []string
func contentToStringArray(c io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
