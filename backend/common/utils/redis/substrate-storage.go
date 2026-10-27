package redis

import (
	"common/utils/strings"
)

var (
	// 锁定基材存储位Key的前缀
	LockedSubstrateStorageLocationKeyPrefix = "locked:substrate_storage_location:"
)

func getLockedSubstrateStorageLocationKey(key string) string {
	return strings.ConcatString(LockedSubstrateStorageLocationKeyPrefix, key)
}

func SetLockedSubstrateStorageLocation(key, value string) {
	Set(getLockedSubstrateStorageLocationKey(key), value, 0)
}

func GetLockedSubstrateStorageLocation(key string) (string, error) {
	return Get(getLockedSubstrateStorageLocationKey(key))
}

func DeleteLockedSubstrateStorageLocation(key string) error {
	return Del(getLockedSubstrateStorageLocationKey(key))
}
