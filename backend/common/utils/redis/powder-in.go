package redis

import (
	"common/utils/strings"
)

var (
	// 锁定粉末入库位Key的前缀
	LockedPowderInLocationKeyPrefix = "locked:powder_in_location:"
)

func getLockedPowderInLocationKey(key string) string {
	return strings.ConcatString(LockedPowderInLocationKeyPrefix, key)
}

func SetLockedPowderInLocation(key, value string) {
	Set(getLockedPowderInLocationKey(key), value, 0)
}

func GetLockedPowderInLocation(key string) (string, error) {
	return Get(getLockedPowderInLocationKey(key))
}

func DeleteLockedPowderInLocation(key string) error {
	return Del(getLockedPowderInLocationKey(key))
}
