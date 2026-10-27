package redis

import (
	"common/utils/strings"
)

var (
	// 锁定基材中转位Key的前缀
	LockedSubstrateTransferKeyPrefix = "locked:substrate_transfer:"
)

func getLockedSubstrateTransferKey(key string) string {
	return strings.ConcatString(LockedSubstrateTransferKeyPrefix, key)
}

func SetLockedSubstrateTransfer(key, value string) {
	Set(getLockedSubstrateTransferKey(key), value, 0)
}

func GetLockedSubstrateTransfer(key string) (string, error) {
	return Get(getLockedSubstrateTransferKey(key))
}

func DeleteLockedSubstrateTransfer(key string) error {
	return Del(getLockedSubstrateTransferKey(key))
}
