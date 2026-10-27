package redis

import (
	"common/utils/strings"
)

var (
	// 锁定零件出库位Key的前缀
	LockedPartOutLocationKeyPrefix = "locked:part_out_location:"
)

func getLockedPartOutLocationKey(key string) string {
	return strings.ConcatString(LockedPartOutLocationKeyPrefix, key)
}

func SetLockedPartOutLocation(key, value string) {
	Set(getLockedPartOutLocationKey(key), value, 0)
}

func GetLockedPartOutLocation(key string) (string, error) {
	return Get(getLockedPartOutLocationKey(key))
}

func DeleteLockedPartOutLocation(key string) error {
	return Del(getLockedPartOutLocationKey(key))
}
