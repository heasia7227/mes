package redis

import (
	"common/utils/strings"
)

var (
	// 锁定倒粉机Key的前缀
	LockedPourOutPowderKeyPrefix = "locked:pour_out_powder:"
)

func getLockedPourOutPowderKey(key string) string {
	return strings.ConcatString(LockedPourOutPowderKeyPrefix, key)
}

func SetLockedPourOutPowder(key, value string) {
	Set(getLockedPourOutPowderKey(key), value, 0)
}

func GetLockedPourOutPowder(key string) (string, error) {
	return Get(getLockedPourOutPowderKey(key))
}

func DeleteLockedPourOutPowder(key string) error {
	return Del(getLockedPourOutPowderKey(key))
}
