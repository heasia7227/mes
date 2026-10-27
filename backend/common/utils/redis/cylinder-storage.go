package redis

import (
	"common/utils/strings"
)

var (
	// 锁定缸体存储位Key的前缀
	LockedCylinderStorageLocationKeyPrefix = "locked:cylinder_storage_location:"
)

func getLockedCylinderStorageLocationKey(key string) string {
	return strings.ConcatString(LockedCylinderStorageLocationKeyPrefix, key)
}

func SetLockedCylinderStorageLocation(key, value string) {
	Set(getLockedCylinderStorageLocationKey(key), value, 0)
}

func GetLockedCylinderStorageLocation(key string) (string, error) {
	return Get(getLockedCylinderStorageLocationKey(key))
}

func DeleteLockedCylinderStorageLocation(key string) error {
	return Del(getLockedCylinderStorageLocationKey(key))
}
