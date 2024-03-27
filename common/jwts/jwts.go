package jwts

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"iotPlatform/common"
	"time"
)

type JwtPayLoad struct {
	Id       uint64 `json:"id"`       //id
	Name     string `json:"username"` // 用户名
	Password string `json:"password"` //用密码签发
}
type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}

// GenToken 创建 Token
func GenToken(user JwtPayLoad) (string, error) {
	accessSecret := common.JwtAccessSecret
	expires := common.JwtAccessExpire
	claim := CustomClaims{
		JwtPayLoad: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expires))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(accessSecret))
}

// ParseToken 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	accessSecret := common.JwtAccessSecret
	//expires := common.JwtAccessExpire
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
