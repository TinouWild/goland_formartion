package basics

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	if len(data) == 0 {
		return "", fmt.Errorf("empty content (file=%v)", filename)
	}

	return string(data), nil
}