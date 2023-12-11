package utils

import (
	"bytes"
	"encoding/gob"
)

// encode takes any Go value and encodes it to a byte slice
func Encode[T any](value T) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(value)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// decode takes a byte slice and decodes it into the specified Go value type
func Decode[T any](encodedData []byte) (T, error) {
	var buffer bytes.Buffer
	buffer.Write(encodedData)
	var decodedValue T
	dec := gob.NewDecoder(&buffer)
	err := dec.Decode(&decodedValue)
	if err != nil {
		return decodedValue, err
	}
	return decodedValue, nil
}
