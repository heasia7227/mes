package redis

import (
	"common/utils/strings"
)

var (
	// 锁定缸体出库位Key的前缀
	LockedCylinderOutLocationKeyPrefix = "locked:cylinder_out_location:"
)

func getLockedCylinderOutLocationKey(key string) string {
	return strings.ConcatString(LockedCylinderOutLocationKeyPrefix, key)
}

func SetLockedCylinderOutLocation(key, value string) {
	Set(getLockedCylinderOutLocationKey(key), value, 0)
}

func GetLockedCylinderOutLocation(key string) (string, error) {
	return Get(getLockedCylinderOutLocationKey(key))
}

func DeleteLockedCylinderOutLocation(key string) error {
	return Del(getLockedCylinderOutLocationKey(key))
}
