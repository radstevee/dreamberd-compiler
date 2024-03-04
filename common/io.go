package common

import (
	"os"
)

func ReadFile(path string) (string, error) {
	fi, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(fi), nil
}
