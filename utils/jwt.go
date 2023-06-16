package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"service/config"
)

// Claims是一些用户信息状态和额外的jwt参数
type Claims struct {
	jwt.StandardClaims
	Id       int    `json:"id"`
	Username string `json:"username"`
}

//GenerateToken 生成token
func GenerateToken(uid int, username string) (string, error) {
	var jwtSecret = []byte(config.ServerGlobalConfig.JwtConfig.JwtSecret) //配置文件中自己配置的
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			Id:     uuid.NewString(),
			Issuer: "hualuerp", //指定发行人
		},
		Id:       uid,
		Username: username,
	}
	// 该方法内部生成签名字符串，再用于获取完整、已签名的token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken @Description: 解析token
func ParseToken(token string) (*Claims, error) {
	var jwtSecret = []byte(config.ServerGlobalConfig.JwtConfig.JwtSecret)
	// 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目结构体都是用指针传递，节省空间
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid { // Valid()验证基于时间的声明
			return claims, nil
		}
	}
	return nil, err
}
