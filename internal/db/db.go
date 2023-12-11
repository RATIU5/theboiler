package db

import (
	"time"

	"go.etcd.io/bbolt"
)

const (
	BUCKET_NAME_CORE = "core"
	BUCKET_KEY_INIT  = "init-name"
)

func OpenDB(path string) (*bbolt.DB, error) {
	return bbolt.Open(path, 0600, &bbolt.Options{Timeout: 1 * time.Second})
}
