package db

import (
	"errors"
	"path/filepath"
	"time"

	"github.com/RATIU5/theboiler/pkg/utils"
	"go.etcd.io/bbolt"
)

const (
	DB_FILENAME    = "data.db"
	DB_BUCKET_CORE = "Core"
)

// Open and return the database connection, if the database doesn't exist, then create it.
//
// Make sure to close the connection after you are done using it.
func OpenDB() (*bbolt.DB, error) {
	path, err := utils.GetAppPath()
	if err != nil {
		return nil, err
	}

	// Create the app directory if it doesn't exist
	err = utils.CreatePath(path)
	if err != nil {
		return nil, err
	}

	path = filepath.Join(path, DB_FILENAME)

	db, err := bbolt.Open(path, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, errors.New("failed to create a connection to the database")
	}

	return db, nil
}
