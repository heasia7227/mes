package redis

import (
	"common/utils/strings"
)

var (
	// 锁定分装设备Key的前缀
	LockedSplitDeviceKeyPrefix = "locked:split_device:"
)

func getLockedSplitDeviceKey(key string) string {
	return strings.ConcatString(LockedSplitDeviceKeyPrefix, key)
}

func SetLockedSplitDevice(key, value string) {
	Set(getLockedSplitDeviceKey(key), value, 0)
}

func GetLockedSplitDevice(key string) (string, error) {
	return Get(getLockedSplitDeviceKey(key))
}

func DeleteLockedSplitDevice(key string) error {
	return Del(getLockedSplitDeviceKey(key))
}
