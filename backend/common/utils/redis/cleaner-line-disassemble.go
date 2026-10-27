package redis

import (
	"common/utils/strings"
)

var (
	// 锁定拆卸平台Key的前缀
	LockedDisassembleKeyPrefix = "locked:disassemble:"
)

func getLockedDisassembleKey(key string) string {
	return strings.ConcatString(LockedDisassembleKeyPrefix, key)
}

func SetLockedDisassemble(key, value string) {
	Set(getLockedDisassembleKey(key), value, 0)
}

func GetLockedDisassemble(key string) (string, error) {
	return Get(getLockedDisassembleKey(key))
}

func DeleteLockedDisassemble(key string) error {
	return Del(getLockedDisassembleKey(key))
}
