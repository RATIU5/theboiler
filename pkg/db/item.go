package db

import (
	"bytes"
	"encoding/gob"

	"github.com/RATIU5/theboiler/pkg/item"
)

func StoreItems(name []byte, items []item.Item) error {
	buffer, err := encodeItem(items)
	if err != nil {
		return err
	}

	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

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

	encData, err := ViewValueInBucket(db, []byte(DB_BUCKET_CORE), name)
	if err != nil {
		return nil, err
	}

	items, err := decodeItem(encData)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func encodeItem(items []item.Item) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(items)
	return buffer.Bytes(), err
}

func decodeItem(encodedData []byte) ([]item.Item, error) {
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
