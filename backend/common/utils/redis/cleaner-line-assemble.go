package redis

import (
	"common/utils/strings"
)

var (
	// 锁定装配平台Key的前缀
	LockedAssembleKeyPrefix = "locked:assemble:"
)

func getLockedAssembleKey(key string) string {
	return strings.ConcatString(LockedAssembleKeyPrefix, key)
}

func SetLockedAssemble(key, value string) {
	Set(getLockedAssembleKey(key), value, 0)
}

func GetLockedAssemble(key string) (string, error) {
	return Get(getLockedAssembleKey(key))
}

func DeleteLockedAssemble(key string) error {
	return Del(getLockedAssembleKey(key))
}
