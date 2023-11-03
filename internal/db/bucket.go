package db

import (
	"errors"
	"fmt"

	"go.etcd.io/bbolt"
)

func CreateBucketIfNotExist(db *bbolt.DB, name string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		return err
	})
}

func SetStringInBucket(db *bbolt.DB, bucket string, key string, value string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return errors.New(fmt.Sprintf("failed to open bucket '%s'", bucket))
		}

		err := b.Put([]byte(key), []byte(value))
		return err
	})
}
