package db

import (
	"fmt"

	"go.etcd.io/bbolt"
)

func CreateBucketIfNotExist(db *bbolt.DB, name []byte) error {
	return db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		return err
	})
}

func SetBytesInBucket(db *bbolt.DB, bucket []byte, key []byte, value []byte) error {
	return db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return fmt.Errorf("failed to open bucket '%s'", bucket)
		}

		err := b.Put(key, value)
		return err
	})
}

func GetBytesInBucket(db *bbolt.DB, bucket []byte, key []byte) ([]byte, error) {
	var value []byte
	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return fmt.Errorf("failed to open bucket '%s'", bucket)
		}

		value = b.Get(key)
		return nil
	})
	if err != nil {
		return []byte(""), err
	}
	return value, nil
}
