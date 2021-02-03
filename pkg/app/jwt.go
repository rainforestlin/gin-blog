package app

import (
	"time"

	"github.com/julianlee107/blogWithGin/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/julianlee107/blogWithGin/global"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    utils.EncodeMD5(appKey),
		AppSecret: utils.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	// 根据claims结构体创建token实例 需要使用jwt.SigningMethodHS256，jwt.SigningMethodES256会报错
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(GetJWTSecret())

}

func ParseToken(token string) (*Claims, error) {
	// 用于解析解析鉴权的声明，方法内部解码校验。最终返回*Token和error
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
