//go:build windows

package appdata

import (
	"os"
	"path/filepath"
)

func dataHome() (string, error) {
	dataHome := os.Getenv("AppData")
	if dataHome == "" {
		home, err := Home()
		if err != nil {
			return "", err
		}

		dataHome = filepath.Join(home, "AppData", "Roaming")
	}

	return dataHome, nil
}

func localDataHome() (string, error) {
	dataHome := os.Getenv("LocalAppData")
	if dataHome == "" {
		home, err := Home()
		if err != nil {
			return "", err
		}

		dataHome = filepath.Join(home, "AppData", "Local")
	}

	return dataHome, nil
}

func configHome() (string, error) {
	configHome := os.Getenv("AppData")
	if configHome == "" {
		home, err := Home()
		if err != nil {
			return "", err
		}

		configHome = filepath.Join(home, "AppData", "Roaming")
	}

	return configHome, nil
}

func localConfigHome() (string, error) {
	configHome := os.Getenv("LocalAppData")
	if configHome == "" {
		home, err := Home()
		if err != nil {
			return "", err
		}

		configHome = filepath.Join(home, "AppData", "Local")
	}

	return configHome, nil
}

func stateHome() (string, error) {
	stateHome := os.TempDir()
	return stateHome, nil
}

func cacheHome() (string, error) {
	cacheHome := os.Getenv("LocalAppData")
	if cacheHome == "" {
		home, err := Home()
		if err != nil {
			return "", err
		}

		cacheHome = filepath.Join(home, "AppData", "Local")
	}

	return cacheHome, nil
}
