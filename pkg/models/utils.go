package models

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

func pointer[K any](val K) *K {
	return &val
}

func rawArrayToJson(value bson.RawValue) (string, error) {
	var bsonArray bson.A
	if err := bson.UnmarshalValue(value.Type, value.Value, &bsonArray); err != nil {
		return "", err
	}

	rawBytes, err := json.Marshal(bsonArray)
	if err != nil {
		return "", err
	}
	return string(rawBytes), nil
}

func rawDocToJson(value bson.RawValue) (string, error) {
	var bsonMap bson.M
	if err := bson.UnmarshalValue(value.Type, value.Value, &bsonMap); err != nil {
		return "", err
	}

	rawBytes, err := json.Marshal(bsonMap)
	if err != nil {
		return "", err
	}
	return string(rawBytes), nil
}
