package redis

import (
	"common/utils/strings"
)

var (
	// 锁定粉末中转位Key的前缀
	LockedPowderTransferKeyPrefix = "locked:powder_transfer:"
)

func getLockedPowderTransferKey(key string) string {
	return strings.ConcatString(LockedPowderTransferKeyPrefix, key)
}

func SetLockedPowderTransfer(key, value string) {
	Set(getLockedPowderTransferKey(key), value, 0)
}

func GetLockedPowderTransfer(key string) (string, error) {
	return Get(getLockedPowderTransferKey(key))
}

func DeleteLockedPowderTransfer(key string) error {
	return Del(getLockedPowderTransferKey(key))
}
