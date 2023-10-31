package db

import (
	"bytes"
	"encoding/gob"

	"github.com/RATIU5/theboiler/pkg/item"
)

func StoreItems(name []byte, items []item.Item) error {
	buffer, err := encode(items)
	if err != nil {
		return err
	}

	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	err = SetValueInBucket(db, []byte(DB_BUCKET_CORE), []byte(name), buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func encode(items []item.Item) (bytes.Buffer, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(items)
	if err != nil {
		return buffer, err
	}
	return buffer, nil
}
