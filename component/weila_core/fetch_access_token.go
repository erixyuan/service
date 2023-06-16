package lanpice_core

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/goutil/strutil"
	"service/cache"
	"service/dto"
	"service/global"
	"service/utils"
	"strconv"
	"time"
)

type FetchAccessTokenResponse struct {
	CoreResp
	Data struct {
		AppToken  string `json:"app_token"`
		ExpiresIn int    `json:"expires_in"`
	} `json:"data"`
}

const (
	lanpiceAppTokenCacheKeyPrefix = "lanpice_app_token_%s"
)

// FetchAccessToken @Description: 获取APP Token
func (s *WeilaService) FetchAccessToken(baseParam dto.BaseParam) string {
	var accessToken string
	key := fmt.Sprintf(lanpiceAppTokenCacheKeyPrefix, s.AppId)
	appAccessToken := cache.RedisClient.Get(key)
	if appAccessToken != nil && appAccessToken.(string) != "" {
		global.GetLogger().InfoByTrace(&baseParam, "从缓存中获取到access token %+v", appAccessToken)
		accessToken = appAccessToken.(string)
	} else {
		timestamp := time.Now().Add(time.Hour * 24).Unix()
		sign := strutil.Md5(strconv.FormatInt(timestamp, 10) + s.AppSecret)

		url := fmt.Sprintf("%s/v1/third/server/app-token?appid=%s&et=%d&sign=%s", s.HostDomain, s.AppId, timestamp, sign)

		global.GetLogger().InfoByTrace(&baseParam, "FetchAccessToken url: %s", url)
		body, err := utils.HTTPGet(url)
		global.GetLogger().InfoByTrace(&baseParam, "FetchAccessToken resp: %s", body)
		if err != nil {
			panic(err)
		}
		var result FetchAccessTokenResponse
		_ = json.Unmarshal(body, &result)
		if result.Error() != nil {
			global.GetLogger().ErrorByTrace(&baseParam, "FetchAccessToken error: %+v", result)
			panic(err)
		}
		accessToken = result.Data.AppToken
		cache.RedisClient.Set(key, accessToken, time.Second*time.Duration(result.Data.ExpiresIn))
	}
	return accessToken
}
