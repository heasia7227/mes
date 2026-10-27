package redis

import "common/utils/strings"

var (
	// 新增清粉线设备OPCUA点位Key的前缀
	OPCUAPointAddCleanerLineDeviceKeyPrefix = "opcua_point:cleaner_line_device:add:"
	// 移除清粉线设备OPCUA点位Key的前缀
	OPCUAPointRemoveCleanerLineDeviceKeyPrefix = "opcua_point:cleaner_line_device:remove:"
)

func getOPCUAPointAddCleanerLineDeviceKey(key string) string {
	return strings.ConcatString(OPCUAPointAddCleanerLineDeviceKeyPrefix, key)
}

func SetOPCUAPointAddCleanerLineDevice(key, value string) {
	Set(getOPCUAPointAddCleanerLineDeviceKey(key), value, 0)
}

func GetOPCUAPointAddCleanerLineDevice(key string) (string, error) {
	return Get(getOPCUAPointAddCleanerLineDeviceKey(key))
}

func DeleteOPCUAPointAddCleanerLineDevice(key string) error {
	return Del(getOPCUAPointAddCleanerLineDeviceKey(key))
}

func GetOPCUAPointAddCleanerLineDeviceKeys() []string {
	keys, err := Keys(OPCUAPointAddCleanerLineDeviceKeyPrefix + "*")
	if err != nil {
		return []string{}
	}
	// 移除前缀
	for i := 0; i < len(keys); i++ {
		keys[i] = strings.TrimPrefix(keys[i], OPCUAPointAddCleanerLineDeviceKeyPrefix)
	}
	return keys
}

func getOPCUAPointRemoveCleanerLineDeviceKey(key string) string {
	return strings.ConcatString(OPCUAPointRemoveCleanerLineDeviceKeyPrefix, key)
}

func SetOPCUAPointRemoveCleanerLineDevice(key, value string) {
	Set(getOPCUAPointRemoveCleanerLineDeviceKey(key), value, 0)
}

func GetOPCUAPointRemoveCleanerLineDevice(key string) (string, error) {
	return Get(getOPCUAPointRemoveCleanerLineDeviceKey(key))
}

func DeleteOPCUAPointRemoveCleanerLineDevice(key string) error {
	return Del(getOPCUAPointRemoveCleanerLineDeviceKey(key))
}

func GetOPCUAPointRemoveCleanerLineDeviceKeys() []string {
	keys, err := Keys(OPCUAPointRemoveCleanerLineDeviceKeyPrefix + "*")
	if err != nil {
		return []string{}
	}
	// 移除前缀
	for i := 0; i < len(keys); i++ {
		keys[i] = strings.TrimPrefix(keys[i], OPCUAPointRemoveCleanerLineDeviceKeyPrefix)
	}
	return keys
}
