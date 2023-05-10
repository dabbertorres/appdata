package appdata

import (
	"os"
	"path/filepath"
	"sync"
)

var (
	setHome sync.Once
	homeDir string
	homeErr error
)

func Home() (string, error) {
	setHome.Do(func() {
		homeDir, homeErr = os.UserHomeDir()
	})
	return homeDir, homeErr
}

func DataHome(appName string) (string, error) {
	dir, err := dataHome()
	if err != nil {
		return "", err
	}

	path := filepath.Join(dir, appName)
	return path, nil
}

func EnsureDataHome(appName string) (string, error) {
	dir, err := DataHome(appName)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(dir, 0o755)
	return dir, err
}

func LocalDataHome(appName string) (string, error) {
	// NOTE: difference is only really relevant for Windows

	dir, err := localDataHome()
	if err != nil {
		return "", err
	}

	path := filepath.Join(dir, appName)
	return path, nil
}

func EnsureLocalDataHome(appName string) (string, error) {
	dir, err := LocalDataHome(appName)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(dir, 0o755)
	return dir, err
}

func ConfigHome(appName string) (string, error) {
	dir, err := configHome()
	if err != nil {
		return "", err
	}

	path := filepath.Join(dir, appName)
	return path, nil
}

func EnsureConfigHome(appName string) (string, error) {
	dir, err := ConfigHome(appName)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(dir, 0o755)
	return dir, err
}

func LocalConfigHome(appName string) (string, error) {
	// NOTE: difference is only really relevant for Windows

	dir, err := localConfigHome()
	if err != nil {
		return "", err
	}

	path := filepath.Join(dir, appName)
	return path, nil
}

func EnsureLocalConfigHome(appName string) (string, error) {
	dir, err := LocalConfigHome(appName)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(dir, 0o755)
	return dir, err
}

func StateHome(appName string) (string, error) {
	dir, err := stateHome()
	if err != nil {
		return "", err
	}

	path := filepath.Join(dir, appName)
	return path, err
}

func EnsureStateHome(appName string) (string, error) {
	dir, err := StateHome(appName)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(dir, 0o755)
	return dir, err
}

func CacheHome(appName string) (string, error) {
	dir, err := cacheHome()
	if err != nil {
		return "", err
	}

	path := filepath.Join(dir, appName)
	return path, nil
}

func EnsureCacheHome(appName string) (string, error) {
	dir, err := CacheHome(appName)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(dir, 0o755)
	return dir, err
}
