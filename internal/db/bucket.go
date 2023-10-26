package db

import (
	"errors"

	"go.etcd.io/bbolt"
)

func DoesBucketExist(db *bbolt.DB, bucketName string) (bool, error) {
	doesBucketExist := false
	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b != nil {
			doesBucketExist = true
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return doesBucketExist, nil
}

func CreateBucketIfNotExist(db *bbolt.DB, bucketName string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})
}

func ViewValueInBucket(db *bbolt.DB, bucketName string, keyName string) (string, error) {
	err := CreateBucketIfNotExist(db, bucketName)
	if err != nil {
		return "", err
	}

	var val string
	err = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return errors.New("bucket not found")
		}

		v := b.Get([]byte(keyName))
		val = string(v)

		return nil
	})

	if err != nil {
		return "", err
	}

	return val, nil
}
