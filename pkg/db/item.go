package db

import (
	"bytes"
	"encoding/gob"

	"github.com/RATIU5/theboiler/pkg/item"
)

func StoreItems(name []byte, items []item.Item) error {
	buffer, err := encodeItems(items)
	if err != nil {
		return err
	}

	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	coreExists := DoesBucketExist(db, DB_BUCKET_CORE)
	if !coreExists {
		CreateBucketIfNotExist(db, DB_BUCKET_CORE)
	}

	err = SetValueInBucket(db, []byte(DB_BUCKET_CORE), []byte(name), buffer)
	if err != nil {
		return err
	}

	return nil
}

func ReadItem(name []byte) ([]item.Item, error) {
	db, err := OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	coreExists := DoesBucketExist(db, DB_BUCKET_CORE)
	if !coreExists {
		CreateBucketIfNotExist(db, DB_BUCKET_CORE)
	}

	encData, err := ViewValueInBucket(db, []byte(DB_BUCKET_CORE), name)
	if err != nil {
		return nil, err
	}

	items, err := decodeItems(encData)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func SetAndReadEncodedValue(items []item.Item) error {
	buffer, err := encodeItems(items)
	if err != nil {
		return err
	}

	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	coreExists := DoesBucketExist(db, DB_BUCKET_CORE)
	if !coreExists {
		CreateBucketIfNotExist(db, DB_BUCKET_CORE)
	}

	err = SetValueInBucket(db, []byte(DB_BUCKET_CORE), []byte("test"), buffer)
	if err != nil {
		return err
	}

	encData, err := ViewValueInBucket(db, []byte(DB_BUCKET_CORE), []byte("test"))
	if err != nil {
		return err
	}

	itms, err := decodeItems(encData)
	if err != nil {
		return err
	}

	for _, itm := range itms {
		itm.Print()
	}

	return nil
}

func encodeItems(items []item.Item) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(items)
	return buffer.Bytes(), err
}

func decodeItems(encodedData []byte) ([]item.Item, error) {
	var buffer bytes.Buffer
	buffer.Write(encodedData)
	var decodedItems []item.Item
	dec := gob.NewDecoder(&buffer)
	err := dec.Decode(&decodedItems)
	if err != nil {
		return nil, err
	}
	return decodedItems, nil
}
