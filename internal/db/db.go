package db

import (
	"errors"
	"time"

	"go.etcd.io/bbolt"
)

const (
	BUCKET_NAME_CORE = "core"
	BUCKET_KEY_INIT  = "init-name"
	BUCKET_KEY_FILES = "files"
)

func OpenDB(path string) (*bbolt.DB, error) {
	return bbolt.Open(path, 0600, &bbolt.Options{Timeout: 1 * time.Second})
}

func CreateCoreBucket(db *bbolt.DB) error {
	err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BUCKET_NAME_CORE))
		return err
	})
	return err
}

func DoesCoreBucketExist(db *bbolt.DB) bool {
	var exists bool
	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_NAME_CORE))
		if b == nil {
			exists = false
			return nil
		}
		exists = true
		return nil
	})
	return exists
}

func DoesBoilerplateExist(db *bbolt.DB, boilerplateName []byte) bool {
	var exists bool
	db.View(func(tx *bbolt.Tx) error {
		bCore := tx.Bucket([]byte(BUCKET_NAME_CORE))
		if bCore == nil {
			exists = false
			return nil
		}

		bName := bCore.Bucket(boilerplateName)
		if bName == nil {
			exists = false
			return nil
		}

		exists = true
		return nil
	})
	return exists
}

func WriteBoilerplate(db *bbolt.DB, boilerplateName []byte, value []byte) error {
	err := db.Update(func(tx *bbolt.Tx) error {
		bCore, err := tx.CreateBucketIfNotExists([]byte(BUCKET_NAME_CORE))
		if err != nil {
			return err
		}

		bName, err := bCore.CreateBucketIfNotExists(boilerplateName)
		if err != nil {
			return err
		}

		err = bName.Put([]byte(BUCKET_KEY_FILES), value)

		return err
	})

	return err
}

func ReadBoilerplate(db *bbolt.DB, bucket []byte) ([]byte, error) {
	var value []byte
	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_NAME_CORE))
		if b == nil {
			return bbolt.ErrBucketNotFound
		}

		bCore := b.Bucket(bucket)
		if bCore == nil {
			return bbolt.ErrBucketNotFound
		}

		value = bCore.Get([]byte(BUCKET_KEY_FILES))
		if value == nil {
			return errors.New("no value found")
		}

		return nil
	})

	return value, err
}
