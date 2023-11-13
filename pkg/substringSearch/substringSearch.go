package substringSearch

import (
	"bytes"
	"os"
)

func Contains(filename string, substring string) (bool, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return false, err
	}
	if bytes.Contains(data, []byte(substring)) {
		return true, nil
	}

	return false, nil
}
