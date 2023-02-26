package util

import (
	"os"
	"path/filepath"
	"strings"
)

func getParentDirectory(levels int) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for i := 0; i < levels; i++ {
		cwd = filepath.Dir(cwd)
	}
	return cwd, nil
}

func WalkDirectory(dir string) error {
	switch dir[0] {
	case '/':
		dir = os.Getenv("PWD") + "/" + dir
	case '~':
		dir = os.Getenv("HOME") + "/" + dir[1:]
	case '.':
		if strings.Contains(dir, "..") {
			// handle directory string with multiple dots
			parts := strings.Split(dir, "/")
			upLevels := 0
			for _, part := range parts[0] {
				if part == '.' {
					upLevels++
				}
			}
			cwd, err := getParentDirectory(upLevels)
			if err != nil {
				return err
			}
			dir = cwd

		} else {
			dir = os.Getenv("PWD") + "/" + dir
		}
	}

	err := os.Chdir(dir)
	if err != nil {
		return err
	}

	return nil
}
