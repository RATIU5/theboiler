package db

import (
	"errors"
	"fmt"
	"log"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	helpers "github.com/RATIU5/theboiler/internal"
	"go.etcd.io/bbolt"
)

const appDBPath = "theboiler/storage/data.db"

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

func getMacPath(homeDir string) string {
	dbPath := filepath.Join(homeDir, fmt.Sprintf("Library/Application Support/%s", appDBPath))
	return dbPath
}

func getLinuxPath(homeDir string) string {
	dbPath := filepath.Join(homeDir, fmt.Sprintf(".local/share/%s", appDBPath))
	return dbPath
}

func getWindowsPath(homeDir string) string {
	dbPath := filepath.Join(homeDir, fmt.Sprintf("AppData/Local/%s", appDBPath))
	return dbPath
}

func OpenDB() {
	path, err := helpers.GetDBPath()
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := bbolt.Open(path, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
