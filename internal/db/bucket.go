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

func DoesBucketExist(db *bbolt.DB, name []byte) bool {
	var exists bool
	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(name))
		if b == nil {
			exists = false
			return nil
		}
		exists = true
		return nil
	})
	return exists
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

func WriteInCore(db *bbolt.DB, bucket []byte, key []byte, value []byte) error {
	err := db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_NAME_CORE))
		if b == nil {
			return fmt.Errorf("failed to open bucket '%s'", BUCKET_NAME_CORE)
		}

		bCore := b.Bucket(bucket)
		if bCore == nil {
			return fmt.Errorf("failed to open bucket '%s'", bucket)
		}

		err := bCore.Put(key, value)

		return err
	})

	return err
}

func ReadFromCore(db *bbolt.DB, bucket []byte, key []byte) ([]byte, error) {
	var value []byte
	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_NAME_CORE))
		if b == nil {
			return fmt.Errorf("failed to open bucket '%s'", BUCKET_NAME_CORE)
		}

		bCore := b.Bucket(bucket)
		if bCore == nil {
			return fmt.Errorf("failed to open bucket '%s'", bucket)
		}

		value = bCore.Get(key)
		if value == nil {
			return fmt.Errorf("failed to get value from key '%s'", key)
		}

		return nil
	})

	return value, err
}
