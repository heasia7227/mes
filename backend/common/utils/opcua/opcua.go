package opcua

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	gopcua "github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
)

// OPCUAClient 封装OPC UA客户端功能
type OPCUAClient struct {
	Client            *gopcua.Client
	EndpointURL       string
	IPAddress         string
	Port              int
	Username          string
	Password          string
	Mutex             sync.RWMutex
	ReconnectInterval time.Duration
	Subscriptions     map[string]context.CancelFunc
}

// NewOPCUAClient 创建一个新的OPC UA客户端实例
func NewOPCUAClient(ipAddress string, port int, username, password string) *OPCUAClient {
	endpointURL := fmt.Sprintf("opc.tcp://%s:%d/freeopcua/server/", ipAddress, port)

	return &OPCUAClient{
		EndpointURL:       endpointURL,
		IPAddress:         ipAddress,
		Port:              port,
		Username:          username,
		Password:          password,
		ReconnectInterval: 1 * time.Second, // 重连间隔1秒
		Subscriptions:     make(map[string]context.CancelFunc),
	}
}

// Connect 连接到OPC UA服务器
func (c *OPCUAClient) Connect() error {
	if c == nil {
		return fmt.Errorf("client is nil, please create instance first")
	}
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	// 创建连接选项 - 移除所有安全模式设置, 让客户端自动协商
	log.Printf("准备连接到OPC UA服务器: %s", c.EndpointURL)
	opts := []gopcua.Option{
		gopcua.SecurityMode(ua.MessageSecurityModeNone),
		gopcua.SecurityPolicy(ua.SecurityPolicyURINone),
		gopcua.AutoReconnect(false),
		// gopcua.ReconnectInterval(c.reconnectInterval),
		gopcua.SessionTimeout(30 * time.Minute), // 延长会话超时
		gopcua.RequestTimeout(10 * time.Second), // 请求超时时间
		gopcua.DialTimeout(10 * time.Second),    // 连接10秒
	}

	// 如果提供了用户名和密码, 则添加认证选项
	if c.Username != "" && c.Password != "" {
		log.Printf("使用用户名密码认证")
		opts = append(opts, gopcua.AuthUsername(c.Username, c.Password))
	}

	// 创建客户端
	log.Printf("创建OPC UA客户端实例")
	client, err := gopcua.NewClient(c.EndpointURL, opts...)
	if err != nil {
		log.Printf("创建OPC UA客户端实例失败: %v", err)
		return err
	}

	// 连接到服务器
	log.Printf("开始连接到服务器")
	if err := client.Connect(context.Background()); err != nil {
		log.Printf("连接失败: %v", err)
		return fmt.Errorf("failed to connect to server: %w", err)
	}

	c.Client = client
	log.Printf("Connected to OPC UA server: %s", c.EndpointURL)
	return nil
}

// Disconnect 断开与OPC UA服务器的连接
func (c *OPCUAClient) Disconnect() error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	// 取消所有订阅
	for id, cancel := range c.Subscriptions {
		cancel()
		delete(c.Subscriptions, id)
		log.Printf("Cancelled subscription %s", id)
	}

	// 断开连接
	if c.Client != nil {
		err := c.Client.Close(context.Background())
		if err != nil {
			return fmt.Errorf("failed to disconnect: %w", err)
		}
		c.Client = nil
		log.Printf("Disconnected from OPC UA server: %s", c.EndpointURL)
	}
	return nil
}

// shouldReconnectOnError 检查是否应该在遇到此错误时进行重新连接
func shouldReconnectOnError(err error) bool {
	// 检查错误消息是否包含需要重新连接的错误类型
	if err == nil {
		return false
	}

	errMsg := err.Error()
	return contains(errMsg, "BadSessionIDInvalid") ||
		contains(errMsg, "0x80250000") || // BadSessionIDInvalid
		contains(errMsg, "BadSessionClosed") ||
		contains(errMsg, "0x80250001") || // BadSessionClosed
		contains(errMsg, "BadSessionTimeout") ||
		contains(errMsg, "0x80250002") || // BadSessionTimeout
		contains(errMsg, "BadServerNotConnected") ||
		contains(errMsg, "0x800D0000") || // BadServerNotConnected
		contains(errMsg, "BadServerHalted") ||
		contains(errMsg, "0x800D0001") || // BadServerHalted
		contains(errMsg, "BadServerShuttingDown") ||
		contains(errMsg, "0x800D0002") || // BadServerShuttingDown
		contains(errMsg, "BadCommunicationError") ||
		contains(errMsg, "0x800F0000") || // BadCommunicationError
		contains(errMsg, "BadConnectionClosed") ||
		contains(errMsg, "0x80130000") || // BadConnectionClosed
		contains(errMsg, "BadTcpConnectionClosed") ||
		contains(errMsg, "0x80860000") || // BadTcpConnectionClosed
		contains(errMsg, "BadTcpSecureChannelClosed") ||
		contains(errMsg, "0x80860002") || // BadTcpSecureChannelClosed
		contains(errMsg, "BadSecureChannelIDInvalid") ||
		contains(errMsg, "0x80150000") || // BadSecureChannelIDInvalid
		contains(errMsg, "BadSecureChannelClosed") ||
		contains(errMsg, "0x80150001") || // BadSecureChannelClosed
		contains(errMsg, "BadTimeout") ||
		contains(errMsg, "0x800C0000") || // BadTimeout
		contains(errMsg, "BadNetworkError") ||
		contains(errMsg, "0x80140000") || // BadNetworkError
		contains(errMsg, "BadConnectionRejected") ||
		contains(errMsg, "0x80130001") || // BadConnectionRejected
		contains(errMsg, "BadServerUnavailable") ||
		contains(errMsg, "0x800D0003") || // BadServerUnavailable
		contains(errMsg, "BadNoCommunication") ||
		contains(errMsg, "0x80120000") || // BadNoCommunication
		contains(errMsg, "BadNotConnected") ||
		contains(errMsg, "0x80130002") || // BadNotConnected
		contains(errMsg, "client not initialized") ||
		contains(errMsg, "connection reset by peer") ||
		contains(errMsg, "connection refused") ||
		contains(errMsg, "dial tcp") ||
		contains(errMsg, "i/o timeout") ||
		contains(errMsg, "context deadline exceeded")
}

// contains 检查字符串s是否包含子字符串sub
func contains(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

// ReadNode 读取单个节点的值
func (c *OPCUAClient) ReadNode(nodeID string) (interface{}, error) {
	// 默认上下文
	return c.ReadNodeWithContext(context.Background(), nodeID)
}

// ReadNodeWithContext 读取单个节点的值 (带上下文)
func (c *OPCUAClient) ReadNodeWithContext(ctx context.Context, nodeID string) (interface{}, error) {
	// 为整个方法添加总超时限制
	totalCtx, totalCancel := context.WithTimeout(ctx, 15*time.Second)
	defer totalCancel()

	c.Mutex.RLock()
	client := c.Client
	c.Mutex.RUnlock()

	if client == nil {
		return nil, fmt.Errorf("client not initialized")
	}

	// 解析节点ID
	node, err := ua.ParseNodeID(nodeID)
	if err != nil {
		return nil, fmt.Errorf("invalid node ID: %w", err)
	}

	// 创建读取请求
	req := &ua.ReadRequest{
		NodesToRead: []*ua.ReadValueID{
			{
				NodeID:      node,
				AttributeID: ua.AttributeIDValue,
			},
		},
	}

	// 发送读取请求
	ctx, cancel := context.WithTimeout(totalCtx, 10*time.Second)
	defer cancel()
	resp, err := client.Read(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("read request failed: %w", err)
	}

	// 检查响应
	if len(resp.Results) == 0 {
		return nil, fmt.Errorf("no results returned")
	}

	result := resp.Results[0]
	if result.Status != ua.StatusOK {
		return nil, fmt.Errorf("read failed with status: %v", result.Status)
	}

	return result.Value.Value(), nil
}

// WriteNode 写入单个节点的值
func (c *OPCUAClient) WriteNode(nodeID string, value interface{}) error {
	// 为整个方法添加总超时限制
	totalCtx, totalCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer totalCancel()

	c.Mutex.RLock()
	client := c.Client
	c.Mutex.RUnlock()

	if client == nil {
		return fmt.Errorf("client not initialized")
	}

	// 解析节点ID
	node, err := ua.ParseNodeID(nodeID)
	if err != nil {
		return fmt.Errorf("invalid node ID: %w", err)
	}

	variant, err := ua.NewVariant(value)
	if err != nil {
		return fmt.Errorf("failed to create variant: %w", err)
	}
	if variant == nil {
		return fmt.Errorf("variant is nil")
	}

	// 创建写入请求
	req := &ua.WriteRequest{
		NodesToWrite: []*ua.WriteValue{
			{
				NodeID:      node,
				AttributeID: ua.AttributeIDValue,
				Value: &ua.DataValue{
					Value:           variant,
					EncodingMask:    ua.DataValueValue | ua.DataValueSourceTimestamp | ua.DataValueServerTimestamp,
					Status:          ua.StatusCode(0), // 必须设为Good
					SourceTimestamp: time.Now(),
					ServerTimestamp: time.Now(),
				},
			},
		},
	}

	// 发送写入请求
	ctx, cancel := context.WithTimeout(totalCtx, 10*time.Second)
	defer cancel()
	resp, err := client.Write(ctx, req)
	log.Printf("写入请求: %v, 响应: %+v, 错误: %v", *req, resp, err)
	if err != nil {
		return fmt.Errorf("write request failed: %w", err)
	}

	// 检查响应
	if len(resp.Results) == 0 {
		return fmt.Errorf("no results returned")
	}

	// 检查响应状态码
	if resp.Results[0] != ua.StatusOK {
		log.Printf("写入响应状态: %v", resp.Results[0])
		return fmt.Errorf("write failed with status: %v", resp.Results[0])
	}

	return nil
}

// IsConnected 检查客户端是否已初始化
func (c *OPCUAClient) IsConnected() bool {
	if c == nil {
		return false
	}
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()
	return c.Client != nil && c.Client.State() == gopcua.Connected
}
