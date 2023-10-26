package utils

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

const (
	APP_PATH = "theboiler/data"
)

// Get the storage path of the DB on MacOS systems
func getMacPath(homeDir string) string {
	dbPath := filepath.Join(homeDir, fmt.Sprintf("Library/Application Support/%s", APP_PATH))
	return dbPath
}

// Get the storage path of the DB on Linux systems
func getLinuxPath(homeDir string) string {
	dbPath := filepath.Join(homeDir, fmt.Sprintf(".local/share/%s", APP_PATH))
	return dbPath
}

// Get the storage path of the DB on Windows systems
func getWindowsPath(homeDir string) string {
	dbPath := filepath.Join(homeDir, fmt.Sprintf("AppData/Local/%s", APP_PATH))
	return dbPath
}

// Get the storage path of the DB for any supported platform
func GetAppPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", errors.New("current user cannot be found")
	}

	homeDir := usr.HomeDir

	switch os := runtime.GOOS; os {
	case "darwin":
		return getMacPath(homeDir), nil
	case "linux":
		return getLinuxPath(homeDir), nil
	case "windows":
		return getWindowsPath(homeDir), nil
	default:
		return "", errors.New("unsupported operating system")
	}
}

// Create a directory path if it doesn't exist.
// The path needs to be valid for the OS type
func CreatePath(path string) error {
	// The permission bits 0755 allow the owner to read, write, and execute while
	// others can read and execute, but not write
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create directory '%s'", path))
	}
	return nil
}
