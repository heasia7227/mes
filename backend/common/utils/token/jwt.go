package token

import (
	"common/utils/conf"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims 自定义载荷, 存放业务信息
type CustomClaims struct {
	jwt.RegisteredClaims
	UserId    string `json:"user_id"`
	UserName  string `json:"user_name"`
	LoginName string `json:"login_name"`
}

// GenerateToken 生成token
func GenerateToken(userId, userName, loginName string) (string, error) {
	// 从配置文件中获取密钥和过期时间
	secretKey := conf.GetStringConf("jwt.secret", "3d-auto-pro")
	accessExpire := conf.GetIntConf("jwt.access_expire", 24)

	now := time.Now()
	claims := CustomClaims{
		UserId:    userId,
		UserName:  userName,
		LoginName: loginName,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),                                              // iat: token签发时间
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(accessExpire) * time.Hour)), // exp: 过期时间
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密钥签名
	return token.SignedString([]byte(secretKey))
}

// ParseToken 解析token, 返回自定义claims
func ParseToken(tokenStr string) (*CustomClaims, error) {
	secretKey := conf.GetStringConf("jwt.secret", "3d-auto-pro")
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 校验签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("非法的签名算法")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token无效")
}

// IsTokenExpired 判断token是否过期
func IsTokenExpired(tokenStr string) (bool, error) {
	claims, err := ParseToken(tokenStr)
	if err != nil {
		// 判断是否是过期错误
		if errors.Is(err, jwt.ErrTokenExpired) {
			return true, nil
		}
		return false, err
	}
	now := time.Now()
	return claims.ExpiresAt.Before(now), nil
}
