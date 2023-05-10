//go:build unix && !darwin

package appdata

import (
	"os"
	"path/filepath"
)

func dataHome() (string, error) {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		home, err := Home()
		if err != nil {
			return "", err
		}

		dataHome = filepath.Join(home, ".local", "share")
	}

	return dataHome, nil
}

func localDataHome() (string, error) { return dataHome() }

func configHome() (string, error) {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		home, err := Home()
		if err != nil {
			return "", err
		}

		configHome = filepath.Join(home, ".config")
	}

	return configHome, nil
}

func localConfigHome() (string, error) { return configHome() }

func stateHome() (string, error) {
	stateHome := os.Getenv("XDG_STATE_HOME")
	if stateHome == "" {
		home, err := Home()
		if err != nil {
			return "", err
		}

		stateHome = filepath.Join(home, ".local", "state")
	}

	return stateHome, nil
}

func cacheHome() (string, error) {
	cacheHome := os.Getenv("XDG_CACHE_HOME")
	if cacheHome == "" {
		home, err := Home()
		if err != nil {
			return "", err
		}

		cacheHome = filepath.Join(home, ".cache")
	}

	return cacheHome, nil
}
