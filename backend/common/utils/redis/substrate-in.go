package redis

import (
	"common/utils/strings"
)

var (
	// 锁定基材入库位Key的前缀
	LockedSubstrateInLocationKeyPrefix = "locked:substrate_in_location:"
)

func getLockedSubstrateInLocationKey(key string) string {
	return strings.ConcatString(LockedSubstrateInLocationKeyPrefix, key)
}

func SetLockedSubstrateInLocation(key, value string) {
	Set(getLockedSubstrateInLocationKey(key), value, 0)
}

func GetLockedSubstrateInLocation(key string) (string, error) {
	return Get(getLockedSubstrateInLocationKey(key))
}

func DeleteLockedSubstrateInLocation(key string) error {
	return Del(getLockedSubstrateInLocationKey(key))
}
