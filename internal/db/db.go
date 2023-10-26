package db

import (
	"path/filepath"
	"time"

	"go.etcd.io/bbolt"
)

const (
	APP_PATH       = "theboiler/data"
	DB_FILENAME    = "data.db"
	DB_BUCKET_CORE = "Core"
)

// Open and return the database connection, if the database doesn't exist, then create it.
//
// Make sure to close the connection after you are done using it.
func OpenDB() (*bbolt.DB, error) {
	path, err := GetAppPath()
	if err != nil {
		return nil, err
	}

	// Create the app directory if it doesn't exist
	err = CreatePath(path)
	if err != nil {
		return nil, err
	}

	path = filepath.Join(path, DB_FILENAME)

	db, err := bbolt.Open(path, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	return db, nil
}
