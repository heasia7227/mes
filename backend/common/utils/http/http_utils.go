package http

import (
	"errors"
	"net"
	"regexp"
	"strings"
)

// ParseUserAgent 解析User-Agent字符串, 返回浏览器和操作系统
func ParseUserAgent(userAgent string) (browser, os string) {
	// 解析操作系统
	osPatterns := map[string]string{
		"Windows":   "Windows",
		"Macintosh": "Mac OS",
		"Linux":     "Linux",
		"Android":   "Android",
		"iPhone":    "iOS",
		"iPad":      "iOS",
	}

	for key, value := range osPatterns {
		if strings.Contains(userAgent, key) {
			os = value
			break
		}
	}

	if os == "" {
		os = "Unknown"
	}

	// 解析浏览器
	browserPatterns := map[string]string{
		"Chrome":  "Chrome",
		"Firefox": "Firefox",
		"Safari":  "Safari",
		"Edge":    "Edge",
		"Opera":   "Opera",
		"MSIE":    "Internet Explorer",
		"Trident": "Internet Explorer",
	}

	for key, value := range browserPatterns {
		if strings.Contains(userAgent, key) {
			browser = value
			break
		}
	}

	if browser == "" {
		browser = "Unknown"
	}

	return
}

// GetLoginLocation 根据IP地址获取登录地点 (这里使用简单的模拟实现, 实际项目中可以使用IP地理位置API)
func GetLoginLocation(ip string) string {
	// 这里使用简单的模拟实现, 实际项目中可以使用IP地理位置API
	// 例如: https://ipstack.com/ 或 https://ipapi.co/

	// 将IPv6本地回环地址转换为IPv4本地回环地址
	if ip == "::1" {
		return "127.0.0.1"
	}

	// 模拟IP地址对应的地理位置
	ipLocationMap := map[string]string{
		"127.0.0.1":      "127.0.0.1",
		"192.168.1.1":    "局域网",
		"192.168.110.70": "公司网络",
	}

	if location, ok := ipLocationMap[ip]; ok {
		return location
	}

	// 对于其他IP地址, 返回IP地址本身
	return ip
}

// IsInternalIP 判断是否为内网IP
func IsInternalIP(ip string) bool {
	// 内网IP范围:
	// 10.0.0.0 ~ 10.255.255.255
	// 172.16.0.0 ~ 172.31.255.255
	// 192.168.0.0 ~ 192.168.255.255

	internalIPRegex := regexp.MustCompile(`^(10\.|172\.(1[6-9]|2[0-9]|3[01])\.|192\.168\.)`)
	return internalIPRegex.MatchString(ip)
}

func IsNetTimeout(err error) bool {
	if err == nil {
		return false
	}

	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return true
	}

	if strings.Contains(err.Error(), "request failed") {
		return true
	}

	return false
}
