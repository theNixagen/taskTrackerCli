package io

import (
	"errors"
	"os"
)

func CreateIfNotExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		os.Create(path)
		return true
	}

	return false
}
