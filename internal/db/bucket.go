package db

import (
	"go.etcd.io/bbolt"
)

func DoesBucketExist(db *bbolt.DB, bucketName string) bool {
	doesBucketExist := false
	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b != nil {
			doesBucketExist = true
		}
		return nil
	})
	return doesBucketExist
}

func CreateBucketIfNotExist(db *bbolt.DB, bucketName string) error {
	if !DoesBucketExist(db, bucketName) {
		db.Update(func(tx *bbolt.Tx) error {
			_, err := tx.CreateBucket([]byte(bucketName))
			if err != nil {
				return err
			}
			return nil
		})
	}
	return nil
}
