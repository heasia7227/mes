package redis

import "common/utils/strings"

var (
	// 新增粉末循环线设备OPCUA点位Key的前缀
	OPCUAPointAddPowderLineDeviceKeyPrefix = "opcua_point:powder_line_device:add:"
	// 移除粉末循环线设备OPCUA点位Key的前缀
	OPCUAPointRemovePowderLineDeviceKeyPrefix = "opcua_point:powder_line_device:remove:"
)

func getOPCUAPointAddPowderLineDeviceKey(key string) string {
	return strings.ConcatString(OPCUAPointAddPowderLineDeviceKeyPrefix, key)
}

func SetOPCUAPointAddPowderLineDevice(key, value string) {
	Set(getOPCUAPointAddPowderLineDeviceKey(key), value, 0)
}

func GetOPCUAPointAddPowderLineDevice(key string) (string, error) {
	return Get(getOPCUAPointAddPowderLineDeviceKey(key))
}

func DeleteOPCUAPointAddPowderLineDevice(key string) error {
	return Del(getOPCUAPointAddPowderLineDeviceKey(key))
}

func GetOPCUAPointAddPowderLineDeviceKeys() []string {
	keys, err := Keys(OPCUAPointAddPowderLineDeviceKeyPrefix + "*")
	if err != nil {
		return []string{}
	}
	// 移除前缀
	for i := 0; i < len(keys); i++ {
		keys[i] = strings.TrimPrefix(keys[i], OPCUAPointAddPowderLineDeviceKeyPrefix)
	}
	return keys
}

func getOPCUAPointRemovePowderLineDeviceKey(key string) string {
	return strings.ConcatString(OPCUAPointRemovePowderLineDeviceKeyPrefix, key)
}

func SetOPCUAPointRemovePowderLineDevice(key, value string) {
	Set(getOPCUAPointRemovePowderLineDeviceKey(key), value, 0)
}

func GetOPCUAPointRemovePowderLineDevice(key string) (string, error) {
	return Get(getOPCUAPointRemovePowderLineDeviceKey(key))
}

func DeleteOPCUAPointRemovePowderLineDevice(key string) error {
	return Del(getOPCUAPointRemovePowderLineDeviceKey(key))
}

func GetOPCUAPointRemovePowderLineDeviceKeys() []string {
	keys, err := Keys(OPCUAPointRemovePowderLineDeviceKeyPrefix + "*")
	if err != nil {
		return []string{}
	}
	// 移除前缀
	for i := 0; i < len(keys); i++ {
		keys[i] = strings.TrimPrefix(keys[i], OPCUAPointRemovePowderLineDeviceKeyPrefix)
	}
	return keys
}
