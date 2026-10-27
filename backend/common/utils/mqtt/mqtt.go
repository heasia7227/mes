package mqtt

import (
	"common/utils/goroutine"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	pahomqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQTTClient 封装MQTT客户端功能
type MQTTClient struct {
	client     pahomqtt.Client
	clientID   string
	brokerAddr string
	username   string
	password   string
	timeout    time.Duration
	onConnect  func(client pahomqtt.Client)
	mu         sync.RWMutex
	topics     map[string]pahomqtt.MessageHandler
}

// MQTTConfig MQTT客户端配置
type MQTTConfig struct {
	BrokerAddr           string
	ClientID             string
	Username             string
	Password             string
	Timeout              time.Duration
	QoS                  byte
	AutoReconnect        bool
	CleanSession         bool
	MaxReconnectInterval time.Duration
}

// NewMQTTClient 创建新的MQTT客户端实例
func NewMQTTClient(config *MQTTConfig) *MQTTClient {
	// 设置默认值
	if config.Timeout <= 0 {
		config.Timeout = 10 * time.Second
	}
	if config.QoS > 2 {
		config.QoS = 1
	}
	if config.ClientID == "" {
		config.ClientID = fmt.Sprintf("mqtt-client-%d", time.Now().UnixNano())
	}

	client := &MQTTClient{
		brokerAddr: config.BrokerAddr,
		clientID:   config.ClientID,
		username:   config.Username,
		password:   config.Password,
		timeout:    config.Timeout,
		topics:     make(map[string]pahomqtt.MessageHandler),
	}

	// 连接选项配置
	opts := pahomqtt.NewClientOptions()
	opts.AddBroker(config.BrokerAddr)
	opts.SetClientID(config.ClientID)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)
	opts.SetCleanSession(config.CleanSession)
	opts.SetAutoReconnect(config.AutoReconnect)
	opts.SetMaxReconnectInterval(config.MaxReconnectInterval)
	opts.SetConnectTimeout(config.Timeout)
	opts.SetKeepAlive(60 * time.Second)

	// 设置回调函数
	opts.SetOnConnectHandler(func(mqttClient pahomqtt.Client) {
		log.Printf("MQTT客户端已连接: %s", config.BrokerAddr)
		// 重新订阅之前的主题
		client.mu.RLock()
		for topic, handler := range client.topics {
			client.client.Subscribe(topic, 1, handler)
			log.Printf("重连后重新订阅主题: %s", topic)
		}
		client.mu.RUnlock()

		if client.onConnect != nil {
			client.onConnect(client.client)
		}
	})

	opts.SetConnectionLostHandler(func(client pahomqtt.Client, err error) {
		log.Printf("MQTT连接丢失: %v", err)
	})

	// 创建MQTT客户端
	client.client = pahomqtt.NewClient(opts)

	return client
}

// Connect 连接到MQTT代理服务器
func (c *MQTTClient) Connect() error {
	// 创建超时通道
	timeout := make(chan bool, 1)
	goroutine.SafeGo(func() {
		time.Sleep(c.timeout)
		timeout <- true
	})

	// 创建结果通道
	result := make(chan error, 1)
	goroutine.SafeGo(func() {
		token := c.client.Connect()
		token.Wait()
		result <- token.Error()
	})

	// 等待结果或超时
	select {
	case err := <-result:
		if err != nil {
			return fmt.Errorf("MQTT连接失败: %w", err)
		}
	case <-timeout:
		return fmt.Errorf("MQTT连接超时")
	}

	log.Printf("MQTT客户端连接成功: %s", c.brokerAddr)
	return nil
}

// Publish 发布消息到指定主题
func (c *MQTTClient) Publish(topic string, qos byte, retained bool, payload interface{}) error {
	// 创建超时通道
	timeout := make(chan bool, 1)
	goroutine.SafeGo(func() {
		time.Sleep(c.timeout)
		timeout <- true
	})

	// 创建结果通道
	result := make(chan error, 1)
	goroutine.SafeGo(func() {
		token := c.client.Publish(topic, qos, retained, payload)
		token.Wait()
		result <- token.Error()
	})

	// 等待结果或超时
	select {
	case err := <-result:
		if err != nil {
			return fmt.Errorf("发布消息失败: %w", err)
		}
	case <-timeout:
		return fmt.Errorf("发布消息超时")
	}

	return nil
}

// PublishJSON 发布JSON格式消息到指定主题
func (c *MQTTClient) PublishJSON(topic string, qos byte, retained bool, data interface{}) error {
	// 将数据序列化为JSON格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("JSON序列化失败: %w", err)
	}

	// 创建超时通道
	timeout := make(chan bool, 1)
	goroutine.SafeGo(func() {
		time.Sleep(c.timeout)
		timeout <- true
	})

	// 创建结果通道
	result := make(chan error, 1)
	goroutine.SafeGo(func() {
		token := c.client.Publish(topic, qos, retained, jsonData)
		token.Wait()
		result <- token.Error()
	})

	// 等待结果或超时
	select {
	case err := <-result:
		if err != nil {
			return fmt.Errorf("发布JSON消息失败: %w", err)
		}
	case <-timeout:
		return fmt.Errorf("发布JSON消息超时")
	}

	log.Printf("已发布JSON消息到主题: %s", topic)
	return nil
}

// Subscribe 订阅主题
func (c *MQTTClient) Subscribe(topic string, qos byte, callback pahomqtt.MessageHandler) error {
	// 创建超时通道
	timeout := make(chan bool, 1)
	goroutine.SafeGo(func() {
		time.Sleep(c.timeout)
		timeout <- true
	})

	// 创建结果通道
	result := make(chan error, 1)
	goroutine.SafeGo(func() {
		token := c.client.Subscribe(topic, qos, callback)
		token.Wait()
		result <- token.Error()
	})

	// 等待结果或超时
	select {
	case err := <-result:
		if err != nil {
			// 订阅失败时删除保存的主题
			c.mu.Lock()
			delete(c.topics, topic)
			c.mu.Unlock()
			return fmt.Errorf("订阅主题失败: %w", err)
		}
		// 保存主题和处理器, 用于重连后重新订阅
		c.mu.Lock()
		c.topics[topic] = callback
		c.mu.Unlock()
	case <-timeout:
		return fmt.Errorf("订阅主题超时")
	}

	log.Printf("已订阅主题: %s", topic)
	return nil
}

// Unsubscribe 取消订阅主题
func (c *MQTTClient) Unsubscribe(topics ...string) error {
	// 创建超时通道
	timeout := make(chan bool, 1)
	goroutine.SafeGo(func() {
		time.Sleep(c.timeout)
		timeout <- true
	})

	// 创建结果通道
	result := make(chan error, 1)
	goroutine.SafeGo(func() {
		token := c.client.Unsubscribe(topics...)
		token.Wait()
		result <- token.Error()
	})

	// 等待结果或超时
	select {
	case err := <-result:
		if err != nil {
			return fmt.Errorf("取消订阅失败: %w", err)
		}
		// 删除保存的主题
		c.mu.Lock()
		for _, topic := range topics {
			delete(c.topics, topic)
			log.Printf("已取消订阅主题: %s", topic)
		}
		c.mu.Unlock()
	case <-timeout:
		return fmt.Errorf("取消订阅超时")
	}

	return nil
}

// Disconnect 断开MQTT连接
func (c *MQTTClient) Disconnect(waitTime time.Duration) {
	c.client.Disconnect(uint(waitTime.Milliseconds()))
	log.Printf("MQTT客户端已断开连接")
}

// IsConnected 检查客户端是否已连接
func (c *MQTTClient) IsConnected() bool {
	return c.client.IsConnected()
}

// SetOnConnectHandler 设置连接成功回调
func (c *MQTTClient) SetOnConnectHandler(callback func(client pahomqtt.Client)) {
	c.onConnect = callback
}

// MQTTError 定义MQTT错误结构
type MQTTError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *MQTTError) Error() string {
	return fmt.Sprintf("MQTT错误 %d: %s", e.Code, e.Message)
}

// ParseJSONMessage 解析JSON格式的消息
func ParseJSONMessage(payload []byte, result interface{}) error {
	if err := json.Unmarshal(payload, result); err != nil {
		return fmt.Errorf("JSON反序列化失败: %w", err)
	}
	return nil
}

// JSONMessageHandler JSON消息处理器生成器
func JSONMessageHandler(handler func(data interface{})) pahomqtt.MessageHandler {
	return func(client pahomqtt.Client, msg pahomqtt.Message) {
		var data interface{}
		if err := ParseJSONMessage(msg.Payload(), &data); err != nil {
			log.Printf("解析JSON消息失败: %v", err)
			msg.Ack()
			return
		}

		log.Printf("收到JSON消息 - 主题: %s", msg.Topic())
		handler(data)
		msg.Ack()
	}
}

// DefaultMessageHandler 默认消息处理器
func DefaultMessageHandler(client pahomqtt.Client, msg pahomqtt.Message) {
	log.Printf("收到消息 - 主题: %s, 负载: %s", msg.Topic(), string(msg.Payload()))
	// 确认消息已接收
	msg.Ack()
}
