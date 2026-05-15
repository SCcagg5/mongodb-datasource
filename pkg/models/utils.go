package models

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

func pointer[K any](val K) *K {
	return &val
}

// rawArrayToJson serializes a BSON array RawValue to a JSON string.
//
// Do not use value.String() here: bsoncore's string formatter can render
// small scientific doubles as values like "2E-05.0", which later fail to
// parse as Extended JSON. Also do not decode directly into bson.A and then
// json.Marshal it: nested primitive.D values are encoded as [{Key,Value}]
// arrays, which changes embedded documents into arrays when tests decode the
// result back through bson.UnmarshalExtJSON.
func rawArrayToJson(value bson.RawValue) (string, error) {
	wrapped := bson.D{{Key: "data", Value: value}}

	extJSON, err := bson.MarshalExtJSON(wrapped, true, false)
	if err != nil {
		return "", err
	}

	var doc bson.M
	if err := bson.UnmarshalExtJSON(extJSON, true, &doc); err != nil {
		return "", err
	}

	rawBytes, err := json.Marshal(doc["data"])
	if err != nil {
		return "", err
	}
	return string(rawBytes), nil
}

// rawDocToJson serializes a BSON embedded-document RawValue to a JSON string.
// See rawArrayToJson for why this goes through MarshalExtJSON instead of
// value.String().
func rawDocToJson(value bson.RawValue) (string, error) {
	extJSON, err := bson.MarshalExtJSON(bson.Raw(value.Value), true, false)
	if err != nil {
		return "", err
	}

	var doc bson.M
	if err := bson.UnmarshalExtJSON(extJSON, true, &doc); err != nil {
		return "", err
	}

	rawBytes, err := json.Marshal(doc)
	if err != nil {
		return "", err
	}
	return string(rawBytes), nil
}
