package bootstrap

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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