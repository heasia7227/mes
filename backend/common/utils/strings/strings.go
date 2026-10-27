package strings

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/google/uuid"
)

// ToString 将任意类型转换为字符串表示
func ToString(v interface{}) string {
	if v == nil {
		return ""
	}

	// 如果是字符串类型, 直接返回, 不进行JSON序列化
	if s, ok := v.(string); ok {
		return s
	}
	// 尝试JSON序列化 (适用于结构体、map等复杂类型)
	if data, err := json.Marshal(v); err == nil {
		return string(data)
	}
	// 如果JSON序列化失败, 使用fmt.Sprintf作为后备方案
	return fmt.Sprintf("%v", v)
}

func ToInt(value interface{}, defaultValue int) int {
	i, err := SafeConvertToInt(value)
	if err != nil {
		return defaultValue
	}
	return i
}

// SafeConvertToInt 安全地将interface{}类型转换为int
func SafeConvertToInt(value interface{}) (int, error) {
	if value == nil {
		return 0, fmt.Errorf("值为nil, 无法转换为int")
	}

	// 尝试不同类型的转换
	switch v := value.(type) {
	case int:
		return v, nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case float32:
		// 检查是否为整数
		if float32(int(v)) == v {
			return int(v), nil
		}
		return 0, fmt.Errorf("float32值 %v 不是整数, 无法安全转换为int", v)
	case float64:
		// 检查是否为整数
		if float64(int(v)) == v {
			return int(v), nil
		}
		return 0, fmt.Errorf("float64值 %v 不是整数, 无法安全转换为int", v)
	default:
		// 尝试将字符串解析为整数
		var i int
		_, err := fmt.Sscanf(ToString(v), "%d", &i)
		if err != nil {
			return 0, fmt.Errorf("字符串 %v 无法解析为int: %v", v, err)
		}
		return i, nil
	}
}

// SafeConvertToInt64 安全地将interface{}类型转换为int64
// 处理可能的float64、string等不同类型输入
func SafeConvertToInt64(value interface{}) (int64, error) {
	if value == nil {
		return 0, fmt.Errorf("值为nil, 无法转换为int64")
	}

	// 尝试不同类型的转换
	switch v := value.(type) {
	case int:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case float32:
		// 检查是否为整数
		if float32(int64(v)) == v {
			return int64(v), nil
		}
		return 0, fmt.Errorf("float32值 %v 不是整数, 无法安全转换为int64", v)
	case float64:
		// 检查是否为整数
		if float64(int64(v)) == v {
			return int64(v), nil
		}
		return 0, fmt.Errorf("float64值 %v 不是整数, 无法安全转换为int64", v)
	case string:
		// 尝试将字符串解析为整数
		var i int64
		_, err := fmt.Sscanf(v, "%d", &i)
		if err != nil {
			return 0, fmt.Errorf("字符串 %v 无法解析为int64: %v", v, err)
		}
		return i, nil
	default:
		return 0, fmt.Errorf("不支持的类型 %T, 无法转换为int64", v)
	}
}

func ToFloat64(value interface{}, defaultValue float64) float64 {
	f, err := SafeConvertToFloat64(value)
	if err != nil {
		return defaultValue
	}
	return f
}

func SafeConvertToFloat64(value interface{}) (float64, error) {
	if value == nil {
		return 0, fmt.Errorf("值为nil, 无法转换为float64")
	}

	// 尝试不同类型的转换
	switch v := value.(type) {
	case int:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float32:
		// 检查是否为整数
		if float32(int(v)) == v {
			return float64(v), nil
		}
		return 0, fmt.Errorf("float32值 %v 不是整数, 无法安全转换为int", v)
	case float64:
		return float64(v), nil
	default:
		// 尝试将字符串解析为整数
		var i float64
		_, err := fmt.Sscanf(ToString(v), "%f", &i)
		if err != nil {
			return 0, fmt.Errorf("字符串 %v 无法解析为float64: %v", v, err)
		}
		return i, nil
	}
}

// ToArray 将字节数组解析为[]map[string]interface{}
func ToArray(data []byte, result *[]map[string]interface{}) error {
	// 使用json.Unmarshal将JSON字节数组解析到数组中
	return json.Unmarshal(data, result)
}

// ToMap 将字节数组解析为map[string]interface{}
func ToMap(data []byte, result *map[string]interface{}) error {
	// 使用json.Unmarshal将JSON字节数组解析到map中
	return json.Unmarshal(data, result)
}

// T 是泛型, 支持传入任何结构体类型
func ToStruct[T any](jsonStr string) (*T, error) {
	var result T
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return &result, fmt.Errorf("JSON解析失败: %w", err)
	}
	return nil, nil
}

// TrimSuffix 移除字符串末尾的指定后缀
func TrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

// Contains 检查字符串是否包含指定子字符串
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// TrimSpace 移除字符串首尾的空格
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// Split 按指定分隔符将字符串分割为字符串片
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// ToUpper 转换为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Join 拼接字符串
func Join(parts []string, sep string) string {
	return strings.Join(parts, sep)
}

// HasPrefix 检查字符串是否以指定前缀开头
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix 检查字符串是否以指定后缀结尾
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex(s, substr string) int {
	return strings.LastIndex(s, substr)
}

// ReplaceAll 替换字符串中的所有子字符串
func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// Repeat 重复字符串
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// 字符串翻转函数 (通用、高效)
func ReverseString(s string) string {
	runes := []rune(s)
	// 双指针翻转
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// GenerateID 生成唯一ID
func GenerateID() string {
	// 剔除横杠, 转大写
	return strings.ToUpper(strings.ReplaceAll(uuid.New().String(), "-", ""))
}

// MD5Hash 对字符串进行MD5加密并返回十六进制字符串
func MD5Hash(str string) string {
	// 创建MD5哈希对象
	hash := md5.New()
	// 写入数据
	hash.Write([]byte(str))
	// 计算哈希值并返回十六进制字符串
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// 字符串构建器池
var stringBuilderPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

// 获取缓冲区的安全方法
func getBuffer() *bytes.Buffer {
	buf := stringBuilderPool.Get().(*bytes.Buffer)
	buf.Reset() // 清空内容但保持容量
	return buf
}

// 归还缓冲区
func putBuffer(buf *bytes.Buffer) {
	stringBuilderPool.Put(buf)
}

// 拼接字符串
func ConcatString(contents ...string) string {
	buf := getBuffer()
	defer putBuffer(buf)

	// 使用 Write 代替 fmt.Sprintf, 减少内存分配
	if len(contents) > 0 {
		for _, content := range contents {
			buf.WriteString(content)
		}
		logContent := buf.String() // 只在最终输出时创建字符串
		return logContent
	}

	return ""
}

// CleanString 清理字符串, 将"-"替换为"_"并移除括号及其内容
// 例如: "C-1-1 (373) " -> "C_1_1"
func CleanString(input string) string {
	// 移除英文括号及其内容
	cleaned := regexp.MustCompile(`\s*\(.*?\)\s*`).ReplaceAllString(input, "")
	// 移除中文括号及其内容
	cleaned = regexp.MustCompile(`\s* (.*?) \s*`).ReplaceAllString(cleaned, "")
	// 将"-"替换为"_"
	cleaned = strings.ReplaceAll(cleaned, "-", "_")
	// 去除首尾空格
	cleaned = strings.TrimSpace(cleaned)
	return cleaned
}

// 给字符串拼接重试次数
func RetryCountConcatString(baseString string) string {
	retryPattern := regexp.MustCompile(`\((\d+)\)[^()]*$`)
	matches := retryPattern.FindStringSubmatch(baseString)
	if len(matches) == 2 {
		// 找到括号中的数字, 将其+1
		retryCount, _ := strconv.Atoi(matches[1])
		newRetryCount := retryCount + 1
		baseString = retryPattern.ReplaceAllString(baseString, fmt.Sprintf("(%d)", newRetryCount))
	} else {
		// 没有括号, 添加 (1) 作为第一次重试
		baseString = fmt.Sprintf("%s(1)", baseString)
	}
	return baseString
}

// CompareVersion 比较两个语义化版本号
// 返回: 1  v1 > v2; 0 v1 == v2; -1 v1 < v2
func CompareVersion(v1, v2 string) int {
	s1 := strings.Split(v1, ".")
	s2 := strings.Split(v2, ".")

	maxLen := len(s1)
	if len(s2) > maxLen {
		maxLen = len(s2)
	}

	for i := 0; i < maxLen; i++ {
		n1 := 0
		if i < len(s1) {
			n1, _ = strconv.Atoi(s1[i])
		}
		n2 := 0
		if i < len(s2) {
			n2, _ = strconv.Atoi(s2[i])
		}

		if n1 > n2 {
			return 1
		} else if n1 < n2 {
			return -1
		}
	}
	return 0
}

// TrimPrefix 移除字符串首部的指定前缀
func TrimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}
