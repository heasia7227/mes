package wlogs

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"common/utils/conf"
	"common/utils/strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 日志写入器缓存, 按 logType 分组
var (
	logWriters = make(map[string]*zap.Logger)
	writersMu  sync.RWMutex
)

// NumberedFileWriter 实现带数字递增的文件写入器
type NumberedFileWriter struct {
	logDir     string
	baseName   string
	maxSize    int64
	mu         sync.Mutex
	file       *os.File
	size       int64
	lastDate   string // 记录上次写入的日期
	skipCaller int    // 跳过调用栈帧数
}

// 设置Log日志的输出目录
func SetOutput() {
	logPath := conf.GetStringConf("logs.file_path", "logs/app")
	baseName := filepath.Base(logPath)
	logFilePath := filepath.Join(logPath, baseName+".log")
	// 创建日志文件
	numberedWriter := NewNumberedFileWriter(logFilePath, 5, 0)
	// 只输出到文件
	log.SetFlags(0)
	log.SetOutput(numberedWriter)
}

// NewNumberedFileWriter 创建新的数字递增文件写入器
func NewNumberedFileWriter(basePath string, maxSizeMB int, skipCaller int) *NumberedFileWriter {
	logDir := filepath.Dir(basePath)
	baseName := filepath.Base(basePath)
	return &NumberedFileWriter{
		logDir:     logDir,
		baseName:   baseName,
		maxSize:    int64(maxSizeMB) * 1024 * 1024,
		skipCaller: skipCaller,
	}
}

// Write 实现 io.Writer 接口
func (w *NumberedFileWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

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

	content := p
	if w.skipCaller > 0 {
		callerInfo := getCallerInfo(w.skipCaller)
		// zap console 编码输出格式为: 时间 \t 级别 \t 消息内容
		// 将调用方信息(如 [file.go:54])插入到日志级别(如 INFO)之后
		firstTab := bytes.IndexByte(p, '\t')
		if firstTab < 0 {
			content = append([]byte(callerInfo+" "), p...)
		} else if rel := bytes.IndexByte(p[firstTab+1:], '\t'); rel >= 0 {
			pos := firstTab + 1 + rel + 1 // 第二个制表符之后的位置
			content = make([]byte, 0, len(p)+len(callerInfo)+1)
			content = append(content, p[:pos]...)
			content = append(content, callerInfo...)
			content = append(content, ' ')
			content = append(content, p[pos:]...)
		} else {
			content = append([]byte(callerInfo+" "), p...)
		}
	}

	now := time.Now().Format("2006-01-02 15:04:05.000")
	// 格式：时间  日志内容
	content = append([]byte(now+" "), content...)

	n, err = w.file.Write(content)
	w.size += int64(n)
	return n, err
}

// getLogTypeDir 获取日志类型目录名 (logDir 的最后一级)
func (w *NumberedFileWriter) getLogTypeDir() string {
	return filepath.Base(w.logDir)
}

// getBaseLogDir 获取基础日志目录 (logDir 去掉最后一级)
func (w *NumberedFileWriter) getBaseLogDir() string {
	return filepath.Dir(w.logDir)
}

// getDateLogDir 获取日期目录路径 (日期在倒数第二级)
func (w *NumberedFileWriter) getDateLogDir(date string) string {
	return filepath.Join(w.getBaseLogDir(), date, w.getLogTypeDir())
}

// open 打开当前日志文件
func (w *NumberedFileWriter) open(date string) error {
	// 确保日期目录存在
	fullLogDir := w.getDateLogDir(date)
	if err := os.MkdirAll(fullLogDir, 0755); err != nil {
		return err
	}
	filePath := filepath.Join(fullLogDir, w.baseName)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	info, err := file.Stat()
	if err != nil {
		file.Close()
		return err
	}
	w.file = file
	w.size = info.Size()
	return nil
}

// rotate 执行日志轮换
func (w *NumberedFileWriter) rotate() {
	if w.file != nil {
		w.file.Close()
		w.file = nil
	}

	// 移动旧文件, 添加数字后缀
	w.moveOldFiles()

	w.size = 0
}

// moveOldFiles 将旧文件重命名为带数字后缀的格式
func (w *NumberedFileWriter) moveOldFiles() {
	if w.lastDate == "" {
		return
	}

	ext := filepath.Ext(w.baseName)
	base := strings.TrimSuffix(w.baseName, ext)
	fullLogDir := w.getDateLogDir(w.lastDate)
	basePath := filepath.Join(fullLogDir, w.baseName)

	// 查找已存在的最大数字后缀
	maxNum := 0
	for i := 1; ; i++ {
		filePath := fmt.Sprintf("%s.%d%s", filepath.Join(fullLogDir, base), i, ext)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			break
		}
		maxNum = i
	}

	// 从最大数字开始, 向后移动文件
	for i := maxNum; i >= 1; i-- {
		src := fmt.Sprintf("%s.%d%s", filepath.Join(fullLogDir, base), i, ext)
		dst := fmt.Sprintf("%s.%d%s", filepath.Join(fullLogDir, base), i+1, ext)
		os.Rename(src, dst)
	}

	// 移动当前文件为 .1
	if _, err := os.Stat(basePath); err == nil {
		dst := fmt.Sprintf("%s.1%s", filepath.Join(fullLogDir, base), ext)
		os.Rename(basePath, dst)
	}
}

// Sync 实现 zapcore.WriteSyncer 接口
func (w *NumberedFileWriter) Sync() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.file != nil {
		return w.file.Sync()
	}
	return nil
}

// 枚举值
const (
	LogTypePrinter                  = "printer"                                  // 打印设备
	LogTypePrinterEnvParams         = "printer-env-params"                       // 打印设备上报环境参数
	LogTypePrinterProgress          = "printer-progress"                         // 打印设备上报进度
	LogTypePrinterAlarmsReport      = "printer-alarms-report"                    // 打印设备报警消息
	LogTypeCleaningLine             = "cleaning-line"                            // 清粉线
	LogTypeCleaningLineAlarm        = "cleaning-line-alarm"                      // 清粉线报警
	LogTypePourOutPowder            = "cleaning-line-pour-out-powder"            // 清粉线倒粉机
	LogTypeDisassemble              = "cleaning-line-disassemble"                // 清粉线拆卸位
	LogTypeAssemble                 = "cleaning-line-assemble"                   // 清粉线装配位
	LogTypeRobotArm                 = "cleaning-line-robot-arm"                  // 清粉线机械臂
	LogTypeCleaningPowder           = "cleaning-line-cleaning-powder"            // 清粉线清粉机
	LogTypeSubstrateTransitPosition = "cleaning-line-substrate-transit-position" // 基材中转位
	LogTypePartTransitPosition      = "cleaning-line-part-transit-position"      // 零件中转位
	LogTypeScanQr                   = "cleaning-line-scan-qr"                    // 扫码位
	LogTypePowderAttract            = "cleaning-line-powder-attract"             // 吸粉位
	LogTypeCylinderStorage          = "location-cylinder-storage"                // 缸体存储位
	LogTypeSubstrateStorage         = "location-substrate-storage"               // 基材存储位
	LogTypePartStorage              = "location-part-storage"                    // 零件存储位
	LogTypeSubstrateIn              = "location-substrate-in"                    // 基材入库位
	LogTypePartOut                  = "location-part-out"                        // 零件出库位
	LogTypeCylinderIn               = "location-cylinder-in"                     // 缸体入库位
	LogTypeCylinderOut              = "location-cylinder-out"                    // 缸体出库位
	LogTypePowderIn                 = "location-powder-in"                       // 粉末入库位
	LogTypePowderLine               = "powder-line"                              // 粉末循环线
	LogTypePowderReceive            = "powder-receive"                           // 粉末循环线接收设备
	LogTypePowderReceiveWeight      = "powder-receive-weight"                    // 粉末循环线接收设备重量
	LogTypePowderSplit              = "powder-split"                             // 粉末循环线分装设备
	LogTypePowderScreening          = "powder-screening"                         // 粉末循环线筛粉机
	LogTypePowderTransitPosition    = "powder-transit-position"                  // 粉末循环线中转位
	LogTypePowderWeightPosition     = "powder-weight-position"                   // 粉末循环线称重位
	LogTypeInOutTask                = "inout-task"                               // 出入库任务监控
	LogTypeAgvTaskAssign            = "agv-task-assign"                          // AGV任务分配
	LogTypeAgvAlarm                 = "agv-alarm"                                // AGV报警
	LogTypeAgvStatus                = "agv-status"                               // AGV状态
	LogTypeMemoryMonitor            = "memory-monitor"                           // 内存监控
	LogTypeAGVCommandRecord         = "agv-command-record"                       // 指令记录
	LogTypeSystemPanic              = "system-panic"                             // 系统异常
	LogTypeSystemCleanup            = "system-cleanup"                           // 系统清理
)

// 写入打印设备日志
// fileName 打印设备编号
// contents 日志内容
func WritePrinterLog(fileName string, contents ...string) {
	contents = append([]string{"[打印设备]"}, contents...)
	WriteLogWithBuffer(LogTypePrinter, fileName, contents...)
}

// 写入打印设备上报环境参数
// fileName: 打印设备编号
// contents: 日志内容
func WritePrinterEnvParamsLog(fileName string, contents ...string) {
	contents = append([]string{"[打印设备环境参数]"}, contents...)
	WriteLogWithBuffer(LogTypePrinterEnvParams, fileName, contents...)
}

// 写入打印设备上报进度
// fileName: 打印设备编号
// contents: 日志内容
func WritePrinterProgressLog(fileName string, contents ...string) {
	contents = append([]string{"[打印设备进度]"}, contents...)
	WriteLogWithBuffer(LogTypePrinterProgress, fileName, contents...)
}

// 写入打印设备报警消息
// fileName: 打印设备编号
// contents: 日志内容
func WritePrinterAlarmsReportLog(fileName string, contents ...string) {
	contents = append([]string{"[打印设备报警]"}, contents...)
	WriteLogWithBuffer(LogTypePrinterAlarmsReport, fileName, contents...)
}

// 写入清粉线日志
// contents: 日志内容
func WriteCleaningLineLog(contents ...string) {
	contents = append([]string{"[清粉线]"}, contents...)
	WriteLogWithBuffer(LogTypeCleaningLine, LogTypeCleaningLine, contents...)
}

// 写入清粉线上报报警
// contents: 日志内容
func WriteCleaningLineAlarmLog(contents ...string) {
	contents = append([]string{"[清粉线报警]"}, contents...)
	WriteLogWithBuffer(LogTypeCleaningLineAlarm, LogTypeCleaningLineAlarm, contents...)
}

// 写入清粉线倒粉机
// fileName: 设备编号
// contents: 日志内容
func WritePourOutPowderLog(fileName string, contents ...string) {
	contents = append([]string{"[倒粉机]"}, contents...)
	WriteLogWithBuffer(LogTypePourOutPowder, fileName, contents...)
}

// 写入清粉线拆卸位
// fileName: 设备编号
// contents: 日志内容
func WriteDisassembleLog(fileName string, contents ...string) {
	contents = append([]string{"[拆卸位]"}, contents...)
	WriteLogWithBuffer(LogTypeDisassemble, fileName, contents...)
}

// 写入清粉线装配位
// fileName: 设备编号
// contents: 日志内容
func WriteAssembleLog(fileName string, contents ...string) {
	contents = append([]string{"[装配位]"}, contents...)
	WriteLogWithBuffer(LogTypeAssemble, fileName, contents...)
}

// 写入清粉线机械臂
// fileName: 设备编号
// contents: 日志内容
func WriteRobotArmLog(fileName string, contents ...string) {
	contents = append([]string{"[机械臂]"}, contents...)
	WriteLogWithBuffer(LogTypeRobotArm, fileName, contents...)
}

// 写入清粉线清粉机
// fileName: 设备编号
// contents: 日志内容
func WriteCleaningPowderLog(fileName string, contents ...string) {
	contents = append([]string{"[清粉机]"}, contents...)
	WriteLogWithBuffer(LogTypeCleaningPowder, fileName, contents...)
}

// 写入清粉线扫码位
// fileName: 设备内位置
// contents: 日志内容
func WriteScanQrLog(fileName string, contents ...string) {
	contents = append([]string{"[扫码位]"}, contents...)
	WriteLogWithBuffer(LogTypeScanQr, fileName, contents...)
}

// 写入清粉线吸粉位
// fileName: 设备内位置
// contents: 日志内容
func WritePowderAttractLog(fileName string, contents ...string) {
	contents = append([]string{"[吸粉位]"}, contents...)
	WriteLogWithBuffer(LogTypePowderAttract, fileName, contents...)
}

// 写入粉末循环线日志
// contents: 日志内容
func WritePowderLineLog(contents ...string) {
	contents = append([]string{"[粉末循环线]"}, contents...)
	WriteLogWithBuffer(LogTypePowderLine, LogTypePowderLine, contents...)
}

// 写入粉末循环线接收设备
// fileName: 设备编号
// contents: 日志内容
func WritePowderReceiveLog(fileName string, contents ...string) {
	contents = append([]string{"[接收设备]"}, contents...)
	WriteLogWithBuffer(LogTypePowderReceive, fileName, contents...)
}

// 写入粉末循环线接收设备重量
// fileName: 设备编号
// contents: 日志内容
func WritePowderReceiveWeightLog(fileName string, contents ...string) {
	contents = append([]string{"[接收设备重量]"}, contents...)
	WriteLogWithBuffer(LogTypePowderReceiveWeight, fileName, contents...)
}

// 写入粉末循环线分装设备
// fileName: 设备编号
// contents: 日志内容
func WritePowderSplitLog(fileName string, contents ...string) {
	contents = append([]string{"[分装设备]"}, contents...)
	WriteLogWithBuffer(LogTypePowderSplit, fileName, contents...)
}

// 写入粉末循环线筛粉机
// fileName: 设备编号
// contents: 日志内容
func WritePowderScreeningLog(fileName string, contents ...string) {
	contents = append([]string{"[筛粉机]"}, contents...)
	WriteLogWithBuffer(LogTypePowderScreening, fileName, contents...)
}

// 写入粉末循环线中转位
// fileName: 库位内位置
// contents: 日志内容
func WritePowderTransitLog(fileName string, contents ...string) {
	contents = append([]string{"[粉末中转位]"}, contents...)
	WriteLogWithBuffer(LogTypePowderTransitPosition, fileName, contents...)
}

// 写入粉末循环线称重位
// fileName: 库位内位置
// contents: 日志内容
func WritePowderWeightLog(fileName string, contents ...string) {
	contents = append([]string{"[粉末称重位]"}, contents...)
	WriteLogWithBuffer(LogTypePowderWeightPosition, fileName, contents...)
}

// 写入基材中转位
// fileName: 库位内位置
// contents: 日志内容
func WriteSubstrateTransitLog(fileName string, contents ...string) {
	contents = append([]string{"[基材中转位]"}, contents...)
	WriteLogWithBuffer(LogTypeSubstrateTransitPosition, fileName, contents...)
}

// 写入零件中转位
// fileName: 库位内位置
// contents: 日志内容
func WritePartTransitLog(fileName string, contents ...string) {
	contents = append([]string{"[零件中转位]"}, contents...)
	WriteLogWithBuffer(LogTypePartTransitPosition, fileName, contents...)
}

// 写入基材入库任务
// contents: 日志内容
func WriteSubstrateInTaskLog(contents ...string) {
	contents = append([]string{"[基材入库任务]"}, contents...)
	WriteLogWithBuffer(LogTypeInOutTask, "substrate_in", contents...)
}

// 写入基材入库单
// contents: 日志内容
func WriteSubstrateInInventoryLog(contents ...string) {
	contents = append([]string{"[基材入库单]"}, contents...)
	WriteLogWithBuffer(LogTypeInOutTask, "substrate_in_inventory", contents...)
}

// 写入缸体入库任务
// contents: 日志内容
func WriteCylinderInTaskLog(contents ...string) {
	contents = append([]string{"[缸体入库任务]"}, contents...)
	WriteLogWithBuffer(LogTypeInOutTask, "cylinder_in", contents...)
}

// 写入缸体入库单
// contents: 日志内容
func WriteCylinderInInventoryLog(contents ...string) {
	contents = append([]string{"[缸体入库单]"}, contents...)
	WriteLogWithBuffer(LogTypeInOutTask, "cylinder_in_inventory", contents...)
}

// 写入缸体出库任务
// contents: 日志内容
func WriteCylinderOutTaskLog(contents ...string) {
	contents = append([]string{"[缸体出库任务]"}, contents...)
	WriteLogWithBuffer(LogTypeInOutTask, "cylinder_out", contents...)
}

// 写入缸体出库单
// contents: 日志内容
func WriteCylinderOutInventoryLog(contents ...string) {
	contents = append([]string{"[缸体出库单]"}, contents...)
	WriteLogWithBuffer(LogTypeInOutTask, "cylinder_out_inventory", contents...)
}

// 写入零件出库任务
// contents: 日志内容
func WritePartOutTaskLog(contents ...string) {
	contents = append([]string{"[零件出库任务]"}, contents...)
	WriteLogWithBuffer(LogTypeInOutTask, "part_out", contents...)
}

// 写入零件出库单
// contents: 日志内容
func WritePartOutInventoryLog(contents ...string) {
	contents = append([]string{"[零件出库单]"}, contents...)
	WriteLogWithBuffer(LogTypeInOutTask, "part_out_inventory", contents...)
}

// 写入缸体存储位
// fileName: 库位编号
// contents: 日志内容
func WriteCylinderStorageLog(fileName string, contents ...string) {
	contents = append([]string{"[缸体存储位]"}, contents...)
	WriteLogWithBuffer(LogTypeCylinderStorage, fileName, contents...)
}

// 写入基材存储位
// fileName: 库位编号
// contents: 日志内容
func WriteSubstrateStorageLog(fileName string, contents ...string) {
	contents = append([]string{"[基材存储位]"}, contents...)
	WriteLogWithBuffer(LogTypeSubstrateStorage, fileName, contents...)
}

// 写入零件存储位
// fileName: 库位编号
// contents: 日志内容
func WritePartStorageLog(fileName string, contents ...string) {
	contents = append([]string{"[零件存储位]"}, contents...)
	WriteLogWithBuffer(LogTypePartStorage, fileName, contents...)
}

// 写入基材入库位
// fileName: 库位编号
// contents: 日志内容
func WriteSubstrateInLog(fileName string, contents ...string) {
	contents = append([]string{"[基材入库位]"}, contents...)
	WriteLogWithBuffer(LogTypeSubstrateIn, fileName, contents...)
}

// 写入零件出库位
// fileName: 库位编号
// contents: 日志内容
func WritePartOutLog(fileName string, contents ...string) {
	contents = append([]string{"[零件出库位]"}, contents...)
	WriteLogWithBuffer(LogTypePartOut, fileName, contents...)
}

// 写入缸体入库位
// fileName: 库位编号
// contents: 日志内容
func WriteCylinderInLog(fileName string, contents ...string) {
	contents = append([]string{"[缸体入库位]"}, contents...)
	WriteLogWithBuffer(LogTypeCylinderIn, fileName, contents...)
}

// 写入缸体出库位
// fileName: 库位编号
// contents: 日志内容
func WriteCylinderOutLog(fileName string, contents ...string) {
	contents = append([]string{"[缸体出库位]"}, contents...)
	WriteLogWithBuffer(LogTypeCylinderOut, fileName, contents...)
}

// 写入零件出库位
// fileName: 库位编号
// contents: 日志内容
func WritePartOutStorageLog(fileName string, contents ...string) {
	contents = append([]string{"[零件出库位]"}, contents...)
	WriteLogWithBuffer(LogTypePartOut, fileName, contents...)
}

// 写入粉末入库位
// fileName: 库位编号
// contents: 日志内容
func WritePowderInLog(fileName string, contents ...string) {
	contents = append([]string{"[粉末入库位]"}, contents...)
	WriteLogWithBuffer(LogTypePowderIn, fileName, contents...)
}

// 写入AGV任务分配
func WriteAgvTaskAssignLog(contents ...string) {
	contents = append([]string{"[AGV任务分配]"}, contents...)
	WriteLogWithBuffer(LogTypeAgvTaskAssign, LogTypeAgvTaskAssign, contents...)
}

// 写入AGV报警
func WriteAgvAlarmLog(contents ...string) {
	contents = append([]string{"[AGV报警]"}, contents...)
	WriteLogWithBuffer(LogTypeAgvAlarm, LogTypeAgvAlarm, contents...)
}

// 写入AGV状态
func WriteAgvStatusLog(contents ...string) {
	contents = append([]string{"[AGV状态]"}, contents...)
	WriteLogWithBuffer(LogTypeAgvStatus, LogTypeAgvStatus, contents...)
}

// 写入内存监控
func WriteMemoryLogStatusLog(contents ...string) {
	contents = append([]string{"[内存监控]"}, contents...)
	WriteLogWithBuffer(LogTypeMemoryMonitor, LogTypeMemoryMonitor, contents...)
}

// 写入AGV指令记录
func WriteAgvCommandRecordLog(contents ...string) {
	contents = append([]string{"[AGV指令记录]"}, contents...)
	WriteLogWithBuffer(LogTypeAGVCommandRecord, LogTypeAGVCommandRecord, contents...)
}

// 写入系统异常
func WriteSystemPanicLog(contents ...string) {
	WriteLogWithBuffer(LogTypeSystemPanic, LogTypeSystemPanic, contents...)
}

// 写入系统清理
func WriteSystemCleanupLog(contents ...string) {
	WriteLogWithBuffer(LogTypeSystemCleanup, LogTypeSystemCleanup, contents...)
}

func getCallerInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	callerInfo := ""
	if ok {
		fileNameOnly := filepath.Base(file)
		callerInfo = fmt.Sprintf("[%s:%d]", fileNameOnly, line)
	}
	return callerInfo
}

// 使用缓存打印日志文件
func WriteLogWithBuffer(logType, fileName string, contents ...string) {
	if fileName == "" {
		fileName = logType
	}

	WriteLogFile(logType, fileName, strings.ConcatString(contents...))
}

// 初始化日志写入器
func initLogger(logType, fileName string) *zap.Logger {
	writersMu.Lock()
	defer writersMu.Unlock()

	if fileName == "" {
		fileName = logType + ".log"
	}

	loggerKey := logType + "/" + fileName
	if logger, ok := logWriters[loggerKey]; ok {
		return logger
	}

	taskLogPath := conf.GetStringConf("logs.task_file_path", "logs/")

	logDir := filepath.Dir(taskLogPath)
	logTypeDir := filepath.Join(logDir, logType)
	logFilePath := filepath.Join(logTypeDir, fileName+".log")

	numberedWriter := NewNumberedFileWriter(logFilePath, 5, 8)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	// 关键：设置OmitKey，移除时间
	encoderConfig.TimeKey = zapcore.OmitKey

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(numberedWriter),
		zapcore.InfoLevel,
	)

	logger := zap.New(core, zap.AddCallerSkip(1))
	logWriters[loggerKey] = logger

	return logger
}

// customTimeEncoder 自定义时间编码器, 保持原有格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// WriteLogFile 写入日志文件,  logType是枚举值
func WriteLogFile(logType, fileName, content string) {
	logger := initLogger(logType, fileName)
	logger.Info(content)
}

// 计算耗时
func CalculateElapsedTime(startTime time.Time) string {
	elapsed := time.Since(startTime)
	return fmt.Sprintf("%.3f 秒", elapsed.Seconds())
}
