package redis

import (
	"common/utils/strings"
)

var (
	// 用户操作 AGV 任务, key 为AGV任务ID, value 为操作状态  1 放弃任务、2 现场手动完成任务
	OperationAgvTaskKeyPrefix = "operation:agv_task:"
	// 处理AGV任务异常映射
	OperationAgvTaskExceptionStepKeyPrefix = "operation:agv_task_exception_step:"
	// 锁定正在分配的AGV任务Key的前缀
	LockedAssigningAgvTaskKeyPrefix = "locked:assigning_agv_task:"
	// 延缓分配AGV任务Key的前缀
	DelayAssigningAgvTaskKeyPrefix = "delay:assigning_agv_task:"
)

func getOperationAgvTaskKey(key string) string {
	return strings.ConcatString(OperationAgvTaskKeyPrefix, key)
}

func SetOperationAgvTask(key, value string) {
	Set(getOperationAgvTaskKey(key), value, 0)
}

func GetOperationAgvTask(key string) (string, error) {
	return Get(getOperationAgvTaskKey(key))
}

func DeleteOperationAgvTask(key string) error {
	return Del(getOperationAgvTaskKey(key))
}

func getOperationAgvTaskExceptionStepKey(key string) string {
	return strings.ConcatString(OperationAgvTaskExceptionStepKeyPrefix, key)
}

func SetOperationAgvTaskExceptionStep(key, value string) {
	Set(getOperationAgvTaskExceptionStepKey(key), value, 0)
}

func GetOperationAgvTaskExceptionStep(key string) (string, error) {
	return Get(getOperationAgvTaskExceptionStepKey(key))
}

func DeleteOperationAgvTaskExceptionStep(key string) error {
	return Del(getOperationAgvTaskExceptionStepKey(key))
}

func getLockedAssigningAgvTaskKey(key string) string {
	return strings.ConcatString(LockedAssigningAgvTaskKeyPrefix, key)
}

func SetLockedAssigningAgvTask(key, value string) {
	Set(getLockedAssigningAgvTaskKey(key), value, 0)
}

func GetLockedAssigningAgvTask(key string) (string, error) {
	return Get(getLockedAssigningAgvTaskKey(key))
}

func DeleteLockedAssigningAgvTask(key string) error {
	return Del(getLockedAssigningAgvTaskKey(key))
}

func DeleteAllLockedAssigningAgvTask() error {
	return Del(getLockedAssigningAgvTaskKey(""))
}

func getDelayAssigningAgvTaskKey(key string) string {
	return strings.ConcatString(DelayAssigningAgvTaskKeyPrefix, key)
}

func SetDelayAssigningAgvTask(key, value string) {
	Set(getDelayAssigningAgvTaskKey(key), value, 0)
}

func GetDelayAssigningAgvTask(key string) (string, error) {
	return Get(getDelayAssigningAgvTaskKey(key))
}

func DeleteDelayAssigningAgvTask(key string) error {
	return Del(getDelayAssigningAgvTaskKey(key))
}
