//go:build darwin

package appdata

import (
	"path/filepath"
)

func dataHome() (string, error) {
	home, err := Home()
	if err != nil {
		return "", err
	}

	dataHome := filepath.Join(home, "Library", "Application Support")
	return dataHome, nil
}

func localDataHome() (string, error) { return dataHome() }

func configHome() (string, error) {
	home, err := Home()
	if err != nil {
		return "", err
	}

	configHome := filepath.Join(home, "Library", "Preferences")
	return configHome, nil
}

func localConfigHome() (string, error) { return configHome() }

func stateHome() (string, error) {
	home, err := Home()
	if err != nil {
		return "", err
	}

	stateHome := filepath.Join(home, "Library", "Caches")
	return stateHome, nil
}

func cacheHome() (string, error) {
	home, err := Home()
	if err != nil {
		return "", err
	}

	cacheHome := filepath.Join(home, "Library", "Application Support")
	return cacheHome, nil
}
