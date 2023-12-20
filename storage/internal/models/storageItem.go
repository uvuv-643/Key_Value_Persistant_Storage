package models

type StorageItem struct {
	Key       string `bson:"key"`
	ValuePath string `bson:"value_path"`
	Extension string `bson:"extension"`
	ExpiresAt uint64 `bson:"expires_at"`
}
