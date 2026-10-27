package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPClient 封装HTTP请求客户端
type HTTPClient struct {
	client  *http.Client
	baseURL string
	headers map[string]string
}

// NewHTTPClient 创建一个新的HTTP客户端实例
func NewHTTPClient(baseURL string, timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
		headers: make(map[string]string),
	}
}

// SetHeader 设置请求头
func (h *HTTPClient) SetHeader(key, value string) {
	h.headers[key] = value
}

// SetHeaders 批量设置请求头
func (h *HTTPClient) SetHeaders(headers map[string]string) {
	for k, v := range headers {
		h.headers[k] = v
	}
}

// filterEmptyValues 过滤掉map中的空值
func filterEmptyValues(data map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range data {
		// 过滤nil值
		if v == nil {
			continue
		}

		// 过滤空字符串
		if str, ok := v.(string); ok && str == "" {
			continue
		}

		// 过滤空切片
		if slice, ok := v.([]interface{}); ok && len(slice) == 0 {
			continue
		}

		// 过滤空map
		if nestedMap, ok := v.(map[string]interface{}); ok {
			// 递归过滤嵌套map
			filteredMap := filterEmptyValues(nestedMap)
			if len(filteredMap) > 0 {
				result[k] = filteredMap
			}
			continue
		}

		// 保留其他值
		result[k] = v
	}
	return result
}

// buildRequest 构建HTTP请求
func (h *HTTPClient) buildRequest(method, url string, body interface{}) (*http.Request, error) {
	var requestBody io.Reader

	// 如果提供了请求体并且不是GET请求, 将其序列化为JSON
	if body != nil && method != http.MethodGet {

		// 在body转字节前, 把值为空的过滤掉
		var filteredBody interface{}
		if dataMap, ok := body.(map[string]interface{}); ok {
			// 如果是map类型, 直接过滤
			filteredBody = filterEmptyValues(dataMap)
		} else {
			// 如果是结构体类型, 先转换为map再过滤
			// 注意: 这种方式只适用于导出字段 (首字母大写)
			bodyBytes, err := json.Marshal(body)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal body for filtering: %w", err)
			}

			var dataMap map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &dataMap); err != nil {
				return nil, fmt.Errorf("failed to unmarshal body for filtering: %w", err)
			}

			filteredBody = filterEmptyValues(dataMap)
		}

		data, err := json.Marshal(filteredBody)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		requestBody = bytes.NewBuffer(data)
	}

	// 构建完整URL
	fullURL := h.baseURL + url

	// 创建请求
	req, err := http.NewRequest(method, fullURL, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// 设置默认Content-Type为application/json
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// 添加自定义请求头
	for k, v := range h.headers {
		req.Header.Set(k, v)
	}

	return req, nil
}

// doRequest 执行HTTP请求并解析响应
func (h *HTTPClient) doRequest(req *http.Request, result interface{}) error {
	// 发送请求
	resp, err := h.client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// 检查HTTP状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	// 如果需要解析响应结果
	if result != nil {
		if err := json.Unmarshal(body, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

// Get 发送GET请求
func (h *HTTPClient) Get(url string, result interface{}) error {
	req, err := h.buildRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	return h.doRequest(req, result)
}

// Post 发送POST请求
func (h *HTTPClient) Post(url string, body interface{}, result interface{}) error {
	req, err := h.buildRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}
	return h.doRequest(req, result)
}

// Put 发送PUT请求
func (h *HTTPClient) Put(url string, body interface{}, result interface{}) error {
	req, err := h.buildRequest(http.MethodPut, url, body)
	if err != nil {
		return err
	}
	return h.doRequest(req, result)
}

// Delete 发送DELETE请求
func (h *HTTPClient) Delete(url string, result interface{}) error {
	req, err := h.buildRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	return h.doRequest(req, result)
}

// Patch 发送PATCH请求
func (h *HTTPClient) Patch(url string, body interface{}, result interface{}) error {
	req, err := h.buildRequest(http.MethodPatch, url, body)
	if err != nil {
		return err
	}
	return h.doRequest(req, result)
}

// PostForm 发送表单数据请求
func (h *HTTPClient) PostForm(url string, data map[string]string, result interface{}) error {
	// 创建表单数据
	formData := &bytes.Buffer{}
	for k, v := range data {
		if formData.Len() > 0 {
			formData.WriteString("&")
		}
		formData.WriteString(fmt.Sprintf("%s=%s", k, v))
	}

	// 构建请求
	req, err := http.NewRequest(http.MethodPost, h.baseURL+url, formData)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// 设置Content-Type为表单
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	// 添加自定义请求头
	for k, v := range h.headers {
		req.Header.Set(k, v)
	}

	return h.doRequest(req, result)
}

// ResponseError 定义响应错误结构
type ResponseError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Details    string `json:"details,omitempty"`
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("HTTP Error %d: %s", e.StatusCode, e.Message)
}

// IsResponseError 检查错误是否为ResponseError类型
func IsResponseError(err error) (*ResponseError, bool) {
	var respErr *ResponseError
	if errors.As(err, &respErr) {
		return respErr, true
	}
	return nil, false
}
