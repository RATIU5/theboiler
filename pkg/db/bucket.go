package db

import (
	"errors"
	"fmt"

	"go.etcd.io/bbolt"
)

// Returns a boolean whether a bucket exists in a database.
func DoesBucketExist(db *bbolt.DB, bucketName string) bool {
	doesBucketExist := false
	_ = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b != nil {
			doesBucketExist = true
		}
		return nil
	})
	return doesBucketExist
}

// Creates the bucket if it does not already exist. It will throw an error if the bucket name is empty or too long.
func CreateBucketIfNotExist(db *bbolt.DB, bucketName string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return errors.New(fmt.Sprintf("could not create bucket, name '%s' is empty or too long", bucketName))
		}
		return nil
	})
}

// Finds a value from the provided key stored in the provided bucket. If the there is no key, an error will be returned.
func ViewValueInBucket(db *bbolt.DB, bucketName string, keyName string) (string, error) {
	var val string
	err := db.View(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return errors.New(fmt.Sprintf("bucket name '%s' is empty or too long", bucketName))
		}
		if b == nil {
			return errors.New(fmt.Sprintf("bucket '%s' created but not found", bucketName))
		}

		v := b.Get([]byte(keyName))
		if v == nil {
			return errors.New(fmt.Sprintf("key '%s' not found in bucket '%s'", keyName, bucketName))
		}
		val = string(v)

		return nil
	})
	if err != nil {
		return "", err
	}

	return val, nil
}

func SetValueInBucket(db *bbolt.DB, bucketName string, keyName string, value string) error {
	err := db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return errors.New(fmt.Sprintf("bucket name '%s' is empty or too long", bucketName))
		}
		if b == nil {
			return errors.New(fmt.Sprintf("bucket '%s' created but not found", bucketName))
		}

		err = b.Put([]byte(keyName), []byte(value))
		if err != nil {
			return errors.New(fmt.Sprintf("key '%s' is too long or value '%s' is too long", keyName, value))
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
