// Package middleware  token校验中间件
package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net/http"
	"time"
	"utils/exception"
)

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 24 * 30 // 设置过期时间

var Secret = []byte(viper.GetString("server.jwtSecret")) // 设置密码，配置文件中读取

// GenerateToken 签发Token
func GenerateToken(userId int64) (string, error) {
	now := time.Now()
	// 创建一个自己的声明
	claims := Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(TokenExpireDuration).Unix(), // 过期时间： 当前时间 + 过期时间
			Issuer:    "admin",                             // 签发人
		},
	}
	// 创建签名对象
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 获取token
	token, err := tokenClaims.SignedString(Secret)

	return token, err
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("invalid token")
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = exception.SUCCESS
		// token 可能在query中也可能在postForm中
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}
		// token不存在
		if token == "" {
			code = exception.RequestERROR
		}

		// 验证token（验证不通过，或者超时）
		claims, err := ParseToken(token)
		if err != nil {
			code = exception.UnAuth
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = exception.TokenTimeOut
		}

		if code != exception.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"StatusCode": code,
				"StatusMsg":  exception.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserId) // 在得到token之后，解析出user_id供全局使用
		// 这样就不用在一些request中都接收token然后都再重新解析成user_id（见前后端交互接口）
		c.Next()
	}
}
