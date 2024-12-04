package basicutil

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func MustFindProjectRoot() string {
	root, err := FindProjectRoot()
	if err != nil {
		panic(err)
	}
	return root
}

func FindProjectRoot() (string, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("could not get current working directory: %w", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(currentPath, "go.mod")); err == nil {
			return currentPath, nil
		}

		nextPath := filepath.Dir(currentPath)
		if nextPath == currentPath {
			return "", errors.New("reached / without finding go.mod")
		}
		currentPath = nextPath
	}
}
