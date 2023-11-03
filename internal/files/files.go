package files

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

const (
	APP_PATH          = "theBoiler/data"
	DATABASE_FILEPATH = "data.db"
)

// Returns true or false if a path exists. If another error
// from calling os.Stat than ErrNotExist, false is returned
func DoesPathExist(path []byte) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	}
	// Default to false
	return false
}

// Creates a file or directory. If the path ends with a
// forward slash, it will be created as a directory. If
// no forward slash is found, a file will be created.
// An error type is returned
func CreatePath(path []byte) error {
	isDir := path[len(path)-1:] == "/"
	if isDir {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
		return nil
	} else {
		file, err := os.Create(path)
		defer file.Close()
		if err != nil {
			return err
		}
		return nil
	}
}

// Retrieve the database file path. Will be located on mac
// in the Library/Application Support folder, windows in the
// AppData/Local folder, and on Linux in the .local/share folder
func GetDatabasePath() []byte {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error: failed to get home directory. reason: %s\n", err)
	}

	var path string
	switch os := runtime.GOOS; os {
	case "darwin":
		path = filepath.Join(home, "Library/Application Support", APP_PATH, DATABASE_FILEPATH)
	case "windows":
		path = filepath.Join(home, "AppData/Local", APP_PATH, DATABASE_FILEPATH)
	case "linux":
		path = filepath.Join(home, ".local/share", APP_PATH, DATABASE_FILEPATH)
	}

	return []byte(path)
}
