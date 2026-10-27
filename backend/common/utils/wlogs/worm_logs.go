package wlogs

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// 北京时间时区，固定 CST +8
var cst = time.FixedZone("CST", 8*3600)

// ormTimestampRe 匹配 ORM 日志行首的 [ORM] + 时间戳, 如 [ORM]2026/09/07 17:50:08
var ormTimestampRe = regexp.MustCompile(`^\[ORM\]\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}`)

type OrmFileWriter struct {
	*NumberedFileWriter
}

var callerPaths = []string{"\\services\\", "\\db\\"}

// NewOrmFileWriter 创建新的数字递增文件写入器
func NewOrmFileWriter(basePath string, maxSizeMB int) *OrmFileWriter {
	fw := NewNumberedFileWriter(basePath, maxSizeMB, 4)
	ow := &OrmFileWriter{
		NumberedFileWriter: fw,
	}

	// 启动兜底协程：每天 0 点强制关闭文件句柄
	go ow.dailyRotateDaemon()

	return ow
}

// dailyRotateDaemon 每日零点强制轮转，不关心中间有没有日志写入
func (w *OrmFileWriter) dailyRotateDaemon() {
	for {
		now := time.Now().In(cst)
		// 计算下一个零点
		nextZero := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, cst)
		timer := time.NewTimer(nextZero.Sub(now))

		<-timer.C
		timer.Stop()

		// 上锁，关闭现有文件，清空日期标记
		w.mu.Lock()
		if w.file != nil {
			_ = w.file.Sync()
			_ = w.file.Close()
			w.file = nil
		}
		w.size = 0
		w.lastDate = "" // 关键：置空，下次Write会用新日期打开
		w.mu.Unlock()
	}
}

// Write 实现 io.Writer 接口
func (w *OrmFileWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	conent := string(p)

	currentDate := time.Now().Format("2006-01-02")

	// 检查日期是否变化
	if w.lastDate != "" && w.lastDate != currentDate {
		w.rotate() // 关闭旧文件, 重置状态
		w.lastDate = currentDate
	}

	// 检查是否需要轮换
	if w.file != nil && w.size+int64(len(p)) > w.maxSize {
		w.rotate()
	}

	// 打开文件 (如果未打开)
	if w.file == nil {
		if err := w.open(currentDate); err != nil {
			return 0, err
		}
		w.lastDate = currentDate
	}

	if strings.Contains(conent, "db.Query") {
		return 0, nil
	}

	callerInfo := ""
	for i := range 20 {
		_, file, line, ok := runtime.Caller(i + 1)
		if ok {
			for _, path := range callerPaths {
				if strings.Contains(filepath.Dir(file), path) {
					// 取出目录层级的最后3级
					dir := filepath.Dir(file)
					parts := strings.Split(filepath.Clean(dir), string(filepath.Separator))
					last3Dir := parts[len(parts)-3:]
					last3File := filepath.Join(last3Dir...)
					baseFile := filepath.Join(last3File, filepath.Base(file))

					callerInfo = fmt.Sprintf("[%s:%d]", baseFile, line)
					break
				}
			}
		}
	}

	if callerInfo != "" {
		// log.go 不打印
		if strings.Contains(callerInfo, "log.go") {
			return 0, nil
		}
		// 原始格式: [ORM]2026/09/07 17:50:08  - ..., 将调用方信息插入到时间戳之后
		if loc := ormTimestampRe.FindStringIndex(conent); loc != nil {
			conent = conent[:loc[1]] + " " + callerInfo + conent[loc[1]:]
		} else {
			conent = fmt.Sprintf("%s %s", callerInfo, conent)
		}
	}

	n, err = w.file.Write([]byte(conent))
	w.size += int64(n)
	return n, err
}
