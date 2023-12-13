package files

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/RATIU5/theboiler/internal/utils"
)

const (
	APP_PATH          = "theBoiler/data"
	DATABASE_FILEPATH = "data.db"
)

// Returns true or false if a path exists. If another error
// from calling os.Stat than ErrNotExist, false is returned
func DoesPathExist(path string) bool {
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
func CreateDirPath(path string) error {
	return os.MkdirAll(path, 0755)
}

// Retrieve the database file path. Will be located on Mac
// in the Library/Application Support folder, Windows in the
// AppData/Local folder, and on Linux in the .local/share folder
func GetDatabasePath() string {
	return filepath.Join(GetApplicationPath(), DATABASE_FILEPATH)
}

// Retrieve the application path. Will be located on Mac
// in the Library/Application Support folder, Windows in the
// AppData/Local folder, and on Linux in the .local/share folder
func GetApplicationPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error: failed to get home directory. reason: %s\n", err)
	}

	var path string
	switch os := runtime.GOOS; os {
	case "darwin":
		path = filepath.Join(home, "Library/Application Support", APP_PATH)
	case "windows":
		path = filepath.Join(home, "AppData/Local", APP_PATH)
	case "linux":
		path = filepath.Join(home, ".local/share", APP_PATH)
	}

	return path
}

// Get the current working directory
func GetWorkingDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error: failed to get working directory. reason: %s\n", err)
	}
	return dir
}

// Get a list of files and folders in a directory
func GetFileList(path string, excludedFiles []string) ([]string, error) {
	var files []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		path = strings.Replace(path, GetWorkingDirectory(), "", 1)
		path = utils.RemoveFirstRune(path)
		if path == "" || path == "." || path == "./" {
			return nil
		}

		for _, excludedFile := range excludedFiles {
			if strings.Contains(path, excludedFile) {
				return nil
			}
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func GetFileContent(path string) (FileContent, error) {
	var fileContent FileContent
	fileContent.Path = path

	file, err := os.Open(path)
	if err != nil {
		return fileContent, err
	}
	defer file.Close()

	// Skip if the file is a directory
	if info, err := file.Stat(); err == nil && info.IsDir() {
		fileContent.IsDir = true
		return fileContent, nil
	} else {
		fileContent.IsDir = false
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContent.Content = append(fileContent.Content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fileContent, err
	}

	return fileContent, nil
}

// Get all the content of many files
func GetFilesContent(paths []string) ([]FileContent, error) {
	var fileContent []FileContent

	for _, path := range paths {
		content, err := GetFileContent(path)
		if err != nil {
			return nil, err
		}
		fileContent = append(fileContent, content)
	}

	return fileContent, nil
}

func CreateDir(path string) error {
	return os.MkdirAll(path, 0755)
}

func CreateFile(path string, content []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range content {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
