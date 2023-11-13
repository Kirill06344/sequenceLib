package substringSearch

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

func Contains(filename string, substring string, readFileByLine bool) (bool, error) {
	if readFileByLine {
		return containsByLine(filename, substring)
	}
	return containsAllFile(filename, substring)
}

func containsByLine(filename string, substring string) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if len(line) == 0 && err != nil {
			if err == io.EOF {
				return false, nil
			}
			return false, nil
		}
		if strings.Contains(line, substring) {
			return true, nil
		}
	}
}

func containsAllFile(filename string, substring string) (bool, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return false, err
	}
	if bytes.Contains(data, []byte(substring)) {
		return true, nil
	}

	return false, nil
}
