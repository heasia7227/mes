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
