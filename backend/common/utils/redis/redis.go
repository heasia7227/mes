package redis

import (
	"common/utils/conf"
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	GlobalRedisClient *redis.Client
	ctx               = context.Background()
)

// InitRedis 初始化Redis客户端连接
// 从环境变量读取Redis连接参数: redis_addr、redis_password、redis_db、redis_pool_size
// 如果参数未配置, 则使用默认值: localhost:6379、空密码、db=0、pool_size=10
func InitRedis() {
	addr := conf.GetStringConf("redis_addr", "127.0.0.1:6379")
	password := conf.GetStringConf("redis_password", "")
	db := conf.GetIntConf("redis_db", 1)
	poolSize := conf.GetIntConf("redis_pool_size", 10)

	GlobalRedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
		PoolSize: poolSize,
	})

	log.Println("正在连接Redis服务器...")
	_, err := GlobalRedisClient.Ping(ctx).Result()
	if err != nil {
		log.Printf("Redis连接失败: %v", err)
		return
	}

	log.Println("Redis客户端已成功初始化并连接")
}

// DisconnectRedis 断开Redis连接
func DisconnectRedis() {
	if GlobalRedisClient != nil {
		if err := GlobalRedisClient.Close(); err != nil {
			log.Printf("Redis断开连接失败: %v", err)
		} else {
			log.Println("Redis连接已断开")
		}
	}
}

// Set 设置键值对, 支持过期时间
// key: 键名
// value: 值 (可以是任意类型, 内部会自动序列化)
// expiration: 过期时间, 类型为time.Duration, 单位为纳秒, 可使用time.Second、time.Minute等常量, 0表示永不过期
// 返回错误信息
func Set(key string, value interface{}, expiration time.Duration) error {
	if GlobalRedisClient == nil {
		return ErrClientNotInitialized
	}
	return GlobalRedisClient.Set(ctx, key, value, expiration).Err()
}

// Get 获取指定键的值
// key: 键名
// 返回值: 键对应的值和错误信息
func Get(key string) (string, error) {
	if GlobalRedisClient == nil {
		return "", ErrClientNotInitialized
	}
	return GlobalRedisClient.Get(ctx, key).Result()
}

// Del 删除指定键
// key: 键名
// 返回错误信息
func Del(key string) error {
	if GlobalRedisClient == nil {
		return ErrClientNotInitialized
	}
	return GlobalRedisClient.Del(ctx, key).Err()
}

// Exists 检查键是否存在
// key: 键名
// 返回值: 键是否存在和错误信息
func Exists(key string) (bool, error) {
	if GlobalRedisClient == nil {
		return false, ErrClientNotInitialized
	}
	result, err := GlobalRedisClient.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

// Incr 对键的值进行原子递增操作
// key: 键名
// 返回值: 递增后的值和错误信息
func Incr(key string) (int64, error) {
	if GlobalRedisClient == nil {
		return 0, ErrClientNotInitialized
	}
	return GlobalRedisClient.Incr(ctx, key).Result()
}

// Decr 对键的值进行原子递减操作
// key: 键名
// 返回值: 递减后的值和错误信息
func Decr(key string) (int64, error) {
	if GlobalRedisClient == nil {
		return 0, ErrClientNotInitialized
	}
	return GlobalRedisClient.Decr(ctx, key).Result()
}

// SetNX 仅当键不存在时设置键值对 (原子操作)
// key: 键名
// value: 值
// expiration: 过期时间
// 返回值: 是否设置成功和错误信息
func SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	if GlobalRedisClient == nil {
		return false, ErrClientNotInitialized
	}
	return GlobalRedisClient.SetNX(ctx, key, value, expiration).Result()
}

// GetSet 设置键的新值并返回旧值
// key: 键名
// value: 新值
// 返回值: 旧值和错误信息
func GetSet(key string, value interface{}) (string, error) {
	if GlobalRedisClient == nil {
		return "", ErrClientNotInitialized
	}
	return GlobalRedisClient.GetSet(ctx, key, value).Result()
}

// TTL 获取键的剩余过期时间
// key: 键名
// 返回值: 剩余过期时间和错误信息
func TTL(key string) (time.Duration, error) {
	if GlobalRedisClient == nil {
		return 0, ErrClientNotInitialized
	}
	return GlobalRedisClient.TTL(ctx, key).Result()
}

// Expire 设置键的过期时间
// key: 键名
// expiration: 过期时间
// 返回值: 是否设置成功和错误信息
func Expire(key string, expiration time.Duration) (bool, error) {
	if GlobalRedisClient == nil {
		return false, ErrClientNotInitialized
	}
	return GlobalRedisClient.Expire(ctx, key, expiration).Result()
}

// HSet 设置Hash类型键的字段值
// key: Hash键名
// values: 字段值对, 格式为key1, value1, key2, value2...
// 返回错误信息
func HSet(key string, values ...interface{}) error {
	if GlobalRedisClient == nil {
		return ErrClientNotInitialized
	}
	return GlobalRedisClient.HSet(ctx, key, values...).Err()
}

// HGet 获取Hash类型键的指定字段值
// key: Hash键名
// field: 字段名
// 返回值: 字段值和错误信息
func HGet(key, field string) (string, error) {
	if GlobalRedisClient == nil {
		return "", ErrClientNotInitialized
	}
	return GlobalRedisClient.HGet(ctx, key, field).Result()
}

// HGetAll 获取Hash类型键的所有字段和值
// key: Hash键名
// 返回值: 字段值映射和错误信息
func HGetAll(key string) (map[string]string, error) {
	if GlobalRedisClient == nil {
		return nil, ErrClientNotInitialized
	}
	return GlobalRedisClient.HGetAll(ctx, key).Result()
}

// HDel 删除Hash类型键的指定字段
// key: Hash键名
// fields: 要删除的字段名列表
// 返回错误信息
func HDel(key string, fields ...string) error {
	if GlobalRedisClient == nil {
		return ErrClientNotInitialized
	}
	return GlobalRedisClient.HDel(ctx, key, fields...).Err()
}

// LPush 在List类型键的左侧插入值
// key: List键名
// values: 要插入的值列表
// 返回错误信息
func LPush(key string, values ...interface{}) error {
	if GlobalRedisClient == nil {
		return ErrClientNotInitialized
	}
	return GlobalRedisClient.LPush(ctx, key, values...).Err()
}

// RPush 在List类型键的右侧插入值
// key: List键名
// values: 要插入的值列表
// 返回错误信息
func RPush(key string, values ...interface{}) error {
	if GlobalRedisClient == nil {
		return ErrClientNotInitialized
	}
	return GlobalRedisClient.RPush(ctx, key, values...).Err()
}

// LPop 从List类型键的左侧弹出一个值
// key: List键名
// 返回值: 弹出的值和错误信息
func LPop(key string) (string, error) {
	if GlobalRedisClient == nil {
		return "", ErrClientNotInitialized
	}
	return GlobalRedisClient.LPop(ctx, key).Result()
}

// RPop 从List类型键的右侧弹出一个值
// key: List键名
// 返回值: 弹出的值和错误信息
func RPop(key string) (string, error) {
	if GlobalRedisClient == nil {
		return "", ErrClientNotInitialized
	}
	return GlobalRedisClient.RPop(ctx, key).Result()
}

// LRange 获取List类型键的指定范围内的元素
// key: List键名
// start: 起始索引 (0表示第一个元素, -1表示最后一个)
// stop: 结束索引 (0表示第一个元素, -1表示最后一个)
// 返回值: 元素列表和错误信息
func LRange(key string, start, stop int64) ([]string, error) {
	if GlobalRedisClient == nil {
		return nil, ErrClientNotInitialized
	}
	return GlobalRedisClient.LRange(ctx, key, start, stop).Result()
}

// SAdd 向Set类型键添加成员
// key: Set键名
// members: 要添加的成员列表
// 返回错误信息
func SAdd(key string, members ...interface{}) error {
	if GlobalRedisClient == nil {
		return ErrClientNotInitialized
	}
	return GlobalRedisClient.SAdd(ctx, key, members...).Err()
}

// SMembers 获取Set类型键的所有成员
// key: Set键名
// 返回值: 成员列表和错误信息
func SMembers(key string) ([]string, error) {
	if GlobalRedisClient == nil {
		return nil, ErrClientNotInitialized
	}
	return GlobalRedisClient.SMembers(ctx, key).Result()
}

// SRem 从Set类型键中移除成员
// key: Set键名
// members: 要移除的成员列表
// 返回错误信息
func SRem(key string, members ...interface{}) error {
	if GlobalRedisClient == nil {
		return ErrClientNotInitialized
	}
	return GlobalRedisClient.SRem(ctx, key, members...).Err()
}

// SIsMember 检查成员是否存在于Set类型键中
// key: Set键名
// member: 要检查的成员
// 返回值: 成员是否存在和错误信息
func SIsMember(key string, member interface{}) (bool, error) {
	if GlobalRedisClient == nil {
		return false, ErrClientNotInitialized
	}
	return GlobalRedisClient.SIsMember(ctx, key, member).Result()
}

// ZAdd 向Sorted Set类型键添加成员及分数
// key: Sorted Set键名
// members: 要添加的成员列表, 每个成员包含分数和值
// 返回错误信息
func ZAdd(key string, members ...*redis.Z) error {
	if GlobalRedisClient == nil {
		return ErrClientNotInitialized
	}
	return GlobalRedisClient.ZAdd(ctx, key, members...).Err()
}

// ZRange 获取Sorted Set类型键指定范围内的成员 (按分数升序)
// key: Sorted Set键名
// start: 起始索引 (0表示第一个元素)
// stop: 结束索引 (-1表示最后一个元素)
// 返回值: 成员列表和错误信息
func ZRange(key string, start, stop int64) ([]string, error) {
	if GlobalRedisClient == nil {
		return nil, ErrClientNotInitialized
	}
	return GlobalRedisClient.ZRange(ctx, key, start, stop).Result()
}

// ZRem 从Sorted Set类型键中移除成员
// key: Sorted Set键名
// members: 要移除的成员列表
// 返回错误信息
func ZRem(key string, members ...interface{}) error {
	if GlobalRedisClient == nil {
		return ErrClientNotInitialized
	}
	return GlobalRedisClient.ZRem(ctx, key, members...).Err()
}

// Keys 查找匹配指定模式的所有键
// pattern: 键名模式, 支持通配符*和?
// 返回值: 匹配的键列表和错误信息
func Keys(pattern string) ([]string, error) {
	if GlobalRedisClient == nil {
		return nil, ErrClientNotInitialized
	}
	return GlobalRedisClient.Keys(ctx, pattern).Result()
}

// FlushDB 清空当前数据库中的所有键
// 返回错误信息
func FlushDB() error {
	if GlobalRedisClient == nil {
		return ErrClientNotInitialized
	}
	return GlobalRedisClient.FlushDB(ctx).Err()
}

// Scan 扫描数据库中的所有键
// 返回值：键列表和错误信息
func Scan(match string) ([]string, error) {
	if GlobalRedisClient == nil {
		return nil, ErrClientNotInitialized
	}
	var keys []string
	var cursor uint64
	for {
		batch, next, err := GlobalRedisClient.Scan(ctx, cursor, match, 100).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, batch...)
		cursor = next
		if cursor == 0 {
			break
		}
	}
	return keys, nil
}

// MGet 获取多个键的值
// keys：要获取的键列表
// 返回值：键值列表（不存在的键对应nil）和错误信息
func MGet(keys ...string) ([]interface{}, error) {
	if GlobalRedisClient == nil {
		return nil, ErrClientNotInitialized
	}
	return GlobalRedisClient.MGet(ctx, keys...).Result()
}

func parseDuration(s string) (time.Duration, error) {
	return time.ParseDuration(s)
}

func parseInt(s string) (int, error) {
	var result int
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, ErrInvalidNumber
		}
		result = result*10 + int(c-'0')
	}
	return result, nil
}

var ErrClientNotInitialized = &redisError{"redis client not initialized"}
var ErrInvalidNumber = &redisError{"invalid number"}

type redisError struct {
	msg string
}

func (e *redisError) Error() string {
	return e.msg
}
