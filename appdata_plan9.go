//go:build plan9

package appdata

import (
	"path/filepath"
)

func dataHome() (string, error) {
	home, err := Home()
	if err != nil {
		return "", err
	}

	dataHome := filepath.Join(home, "lib")
	return dataHome, nil
}

func localDataHome() (string, error) { return dataHome() }

func configHome() (string, error) {
	home, err := Home()
	if err != nil {
		return "", err
	}

	configHome := filepath.Join(home, "lib")
	return configHome, nil
}

func localConfigHome() (string, error) { return configHome() }

func stateHome() (string, error) {
	home, err := Home()
	if err != nil {
		return "", err
	}

	stateHome := filepath.Join(home, "lib", "cache")
	return stateHome, nil
}

func cacheHome() (string, error) {
	home, err := Home()
	if err != nil {
		return "", err
	}

	cacheHome := filepath.Join(home, "lib", "cache")
	return cacheHome, nil
}
