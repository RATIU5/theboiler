package db

import (
	"errors"
	"fmt"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	"go.etcd.io/bbolt"
)

const appDBPath = "theboiler/storage/data.db"

// Get the storage path of the DB on MacOS systems
func getMacPath(homeDir string) string {
	dbPath := filepath.Join(homeDir, fmt.Sprintf("Library/Application Support/%s", appDBPath))
	return dbPath
}

// Get the storage path of the DB on Linux systems
func getLinuxPath(homeDir string) string {
	dbPath := filepath.Join(homeDir, fmt.Sprintf(".local/share/%s", appDBPath))
	return dbPath
}

// Get the storage path of the DB on Windows systems
func getWindowsPath(homeDir string) string {
	dbPath := filepath.Join(homeDir, fmt.Sprintf("AppData/Local/%s", appDBPath))
	return dbPath
}

// Get the storage path of the DB for any supported platform
func GetDBPath() (string, error) {
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

// Open and return the database connection, if the database doesn't exist, then create it
func OpenDB() (*bbolt.DB, error) {
	path, err := GetDBPath()
	if err != nil {
		return nil, err
	}

	db, err := bbolt.Open(path, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	return db, nil
}
