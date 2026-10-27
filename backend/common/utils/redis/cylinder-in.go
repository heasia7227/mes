package redis

import (
	"common/utils/strings"
)

var (
	// 锁定缸体入库位Key的前缀
	LockedCylinderInLocationKeyPrefix = "locked:cylinder_in_location:"
)

func getLockedCylinderInLocationKey(key string) string {
	return strings.ConcatString(LockedCylinderInLocationKeyPrefix, key)
}

func SetLockedCylinderInLocation(key, value string) {
	Set(getLockedCylinderInLocationKey(key), value, 0)
}

func GetLockedCylinderInLocation(key string) (string, error) {
	return Get(getLockedCylinderInLocationKey(key))
}

func DeleteLockedCylinderInLocation(key string) error {
	return Del(getLockedCylinderInLocationKey(key))
}
