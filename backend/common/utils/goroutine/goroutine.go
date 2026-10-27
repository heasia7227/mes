package goroutine

import (
	"common/utils/wlogs"
	"fmt"
	"runtime"
)

// SafeGo 安全启动goroutine，内置recover + 堆栈捕获
// f: 业务任务函数
func SafeGo(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				// 打印堆栈，方便定位问题
				buf := make([]byte, 1024*8)
				buf = buf[:runtime.Stack(buf, false)]
				wlogs.WriteSystemPanicLog(fmt.Sprintf("panic: %v\nstack: %s\n", err, buf))
			}
		}()

		f()
	}()
}
