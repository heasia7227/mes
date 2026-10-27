package redis

import (
	"common/utils/strings"
)

var (
	// 锁定零件中转位Key的前缀
	LockedPartTransferLocationKeyPrefix = "locked:part_transfer_location:"
)

func getLockedPartTransferLocationKey(key string) string {
	return strings.ConcatString(LockedPartTransferLocationKeyPrefix, key)
}

func SetLockedPartTransferLocation(key, value string) {
	Set(getLockedPartTransferLocationKey(key), value, 0)
}

func GetLockedPartTransferLocation(key string) (string, error) {
	return Get(getLockedPartTransferLocationKey(key))
}

func DeleteLockedPartTransferLocation(key string) error {
	return Del(getLockedPartTransferLocationKey(key))
}
