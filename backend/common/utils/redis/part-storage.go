package redis

import (
	"common/utils/strings"
)

var (
	// 锁定零件存储位Key的前缀
	LockedPartStorageLocationKeyPrefix = "locked:part_storage_location:"
)

func getLockedPartStorageLocationKey(key string) string {
	return strings.ConcatString(LockedPartStorageLocationKeyPrefix, key)
}

func SetLockedPartStorageLocation(key, value string) {
	Set(getLockedPartStorageLocationKey(key), value, 0)
}

func GetLockedPartStorageLocation(key string) (string, error) {
	return Get(getLockedPartStorageLocationKey(key))
}

func DeleteLockedPartStorageLocation(key string) error {
	return Del(getLockedPartStorageLocationKey(key))
}
