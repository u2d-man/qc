package cli

import (
	"fmt"
	"io/ioutil"
	"os"
)

type SQL struct {
	fields []string
}

func readFile(fn string) (string, error) {
	f, err := os.Open(fn)
	if err != nil {
		return "", fmt.Errorf("File open error: %v", err)
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("File read error: %v", err)
	}

	return string(b), nil
}
