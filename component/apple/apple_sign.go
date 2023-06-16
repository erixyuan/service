package apple

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

const (
	GetApplePublicKeys = "https://appleid.apple.com/auth/keys"
	AppleUrl           = "https://appleid.apple.com"
	ClientId           = "com.dinggeqiu.blindbox"
)

type JwtClaims struct {
	CHash          string `json:"c_hash"`
	Email          string `json:"email"`
	EmailVerified  string `json:"email_verified"`
	AuthTime       int    `json:"auth_time"`
	NonceSupported bool   `json:"nonce_supported"`
	jwt.StandardClaims
}

type JwtHeader struct {
	Kid string `json:"kid"`
	Alg string `json:"alg"`
}

type JwtKeys struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}

// VerifyIdentityToken @Description: 认证客户端传递过来的token是否有效
func VerifyIdentityToken(cliToken string, cliUserID string) error {
	cliTokenArr := strings.Split(cliToken, ".")
	if len(cliTokenArr) < 3 {
		return errors.New("cliToken split err")
	}
	cliHeader, err := jwt.DecodeSegment(cliTokenArr[0])
	if err != nil {
		return err
	}
	var jHeader JwtHeader
	err = json.Unmarshal(cliHeader, &jHeader)
	if err != nil {
		return err
	}
	//  校验pubKey及token
	token, err := jwt.ParseWithClaims(cliToken, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return GetRSAPublicKey(jHeader.Kid), nil
	})

	if err != nil {
		return err
	}
	//  信息验证
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		if claims.StandardClaims.Issuer != AppleUrl || claims.StandardClaims.Audience != ClientId || claims.StandardClaims.Subject != cliUserID {
			return errors.New("verify token info fail, info is not match")
		}
		return nil
	}
	return errors.New("token claims parse fail")
}

func GetRSAPublicKey(kid string) *rsa.PublicKey {
	response, err := http.Get(GetApplePublicKeys)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	var jKeys map[string][]JwtKeys
	err = json.Unmarshal(body, &jKeys)
	if err != nil {
		return nil
	}
	var pubKey rsa.PublicKey

	for _, data := range jKeys {
		for _, val := range data {
			if val.Kid == kid {
				nByte, _ := base64.RawURLEncoding.DecodeString(val.N)
				nData := new(big.Int).SetBytes(nByte)

				eByte, _ := base64.RawURLEncoding.DecodeString(val.E)
				eData := new(big.Int).SetBytes(eByte)

				pubKey.N = nData
				pubKey.E = int(eData.Uint64())
				break
			}
		}
	}

	if pubKey.E <= 0 {
		return nil
	}
	return &pubKey
}
