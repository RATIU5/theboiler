package db

import (
	"os/user"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGetDBPath(t *testing.T) {
	path, err := GetAppPath()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if path == "" {
		t.Fatalf("expected a non-empty path")
	}
}

func TestGetMacPath(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("skipping test on non-mac system")
	}

	usr, err := user.Current()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	path := getMacPath(usr.HomeDir)
	expectedPath := filepath.Join(usr.HomeDir, "Library/Application Support", APP_PATH)

	if path != expectedPath {
		t.Fatalf("expected %s but got %s", expectedPath, path)
	}
}

func TestGetLinuxPath(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("skipping test on non-linux system")
	}

	usr, err := user.Current()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	path := getLinuxPath(usr.HomeDir)
	expectedPath := filepath.Join(usr.HomeDir, ".local/share", APP_PATH)

	if path != expectedPath {
		t.Fatalf("expected %s but got %s", expectedPath, path)
	}
}

func TestGetWindowsPath(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("skipping test on non-windows system")
	}

	usr, err := user.Current()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	path := getWindowsPath(usr.HomeDir)
	expectedPath := filepath.Join(usr.HomeDir, "AppData/Local", APP_PATH)

	if path != expectedPath {
		t.Fatalf("expected %s but got %s", expectedPath, path)
	}
}
