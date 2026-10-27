package redis

import (
	"common/utils/strings"
)

var (
	// 锁定接收设备Key的前缀
	LockedReceiveDeviceKeyPrefix = "locked:receive_device:"
)

func getLockedReceiveDeviceKey(key string) string {
	return strings.ConcatString(LockedReceiveDeviceKeyPrefix, key)
}

func SetLockedReceiveDevice(key, value string) {
	Set(getLockedReceiveDeviceKey(key), value, 0)
}

func GetLockedReceiveDevice(key string) (string, error) {
	return Get(getLockedReceiveDeviceKey(key))
}

func DeleteLockedReceiveDevice(key string) error {
	return Del(getLockedReceiveDeviceKey(key))
}
