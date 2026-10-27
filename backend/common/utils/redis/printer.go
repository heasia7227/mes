package redis

import (
	"common/utils/strings"
)

var (
	// 锁定打印设备Key的前缀
	LockedPrinterKeyPrefix = "locked:printer:"
	// 锁定可打印设备Key的前缀
	LockedPrintablePrinterKeyPrefix = "locked:printable_printer:"
	// 锁定正在分配的打印设备Key的前缀
	LockedAssigningPrinterKeyPrefix = "locked:assigning_printer:"
	// 锁定正在分配的生产任务Key的前缀
	LockedAssigningTaskKeyPrefix = "locked:assigning_task:"
)

func getLockedPrinterKey(key string) string {
	return strings.ConcatString(LockedPrinterKeyPrefix, key)
}

func SetLockedPrinter(key, value string) {
	Set(getLockedPrinterKey(key), value, 0)
}

func GetLockedPrinter(key string) (string, error) {
	return Get(getLockedPrinterKey(key))
}

func DeleteLockedPrinter(key string) error {
	return Del(getLockedPrinterKey(key))
}

func getLockedPrintablePrinterKey(key string) string {
	return strings.ConcatString(LockedPrintablePrinterKeyPrefix, key)
}

func SetLockedPrintablePrinter(key, value string) {
	Set(getLockedPrintablePrinterKey(key), value, 0)
}

func GetLockedPrintablePrinter(key string) (string, error) {
	return Get(getLockedPrintablePrinterKey(key))
}

func DeleteLockedPrintablePrinter(key string) error {
	return Del(getLockedPrintablePrinterKey(key))
}

func GetLockedPrintablePrinterKeys() []string {
	keys, err := Keys(LockedPrintablePrinterKeyPrefix + "*")
	if err != nil {
		return []string{}
	}
	// 移除前缀
	for i := 0; i < len(keys); i++ {
		keys[i] = strings.TrimPrefix(keys[i], LockedPrintablePrinterKeyPrefix)
	}
	return keys
}

func getLockedAssigningPrinterKey(key string) string {
	return strings.ConcatString(LockedAssigningPrinterKeyPrefix, key)
}

func SetLockedAssigningPrinter(key, value string) {
	Set(getLockedAssigningPrinterKey(key), value, 0)
}

func GetLockedAssigningPrinter(key string) (string, error) {
	return Get(getLockedAssigningPrinterKey(key))
}

func DeleteLockedAssigningPrinter(key string) error {
	return Del(getLockedAssigningPrinterKey(key))
}

func getLockedAssigningTaskKey(key string) string {
	return strings.ConcatString(LockedAssigningTaskKeyPrefix, key)
}

func SetLockedAssigningTask(key, value string) {
	Set(getLockedAssigningTaskKey(key), value, 0)
}

func GetLockedAssigningTask(key string) (string, error) {
	return Get(getLockedAssigningTaskKey(key))
}

func DeleteLockedAssigningTask(key string) error {
	return Del(getLockedAssigningTaskKey(key))
}

func GetLockedAssigningTaskKeys() []string {
	keys, err := Keys(LockedAssigningTaskKeyPrefix + "*")
	if err != nil {
		return []string{}
	}
	// 移除前缀
	for i := 0; i < len(keys); i++ {
		keys[i] = strings.TrimPrefix(keys[i], LockedAssigningTaskKeyPrefix)
	}
	return keys
}
