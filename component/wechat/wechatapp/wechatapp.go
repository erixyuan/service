package wechatapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"service/cache"
	"service/global"
	"service/static"
	"service/utils"
	"time"
)

const (
	codeAuthUrl = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	userInfoUrl = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s"
)

type ErrorCode struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type CodeAuthResult struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
}

type UserInfoResult struct {
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

var (
	WxAppClient *WeChatAppClient
)

type WeChatAppClient struct {
	AppId     string
	AppSecret string
}

func InitWeChatAppClient(appId string, appSecret string) {
	WxAppClient = &WeChatAppClient{
		AppId:     appId,
		AppSecret: appSecret,
	}
}

// CodeAuth @Description: 通过 code 获取 access_token
func (c *WeChatAppClient) CodeAuth(code string) (*CodeAuthResult, error) {
	url := fmt.Sprintf(codeAuthUrl, c.AppId, c.AppSecret, code)
	body, err := utils.HTTPGet(url)
	if err != nil {
		return nil, err
	}
	var result CodeAuthResult
	_ = json.Unmarshal(body, &result)
	if result.AccessToken == "" {
		global.GetLogger().Errorf("wechat app code auth is fail: %s", string(body))
		return nil, errors.New("wechat app code auth is error by code:" + code)
	}
	return &result, nil
}

// FetchUserInfo @Description: 获取用户信息
func (c *WeChatAppClient) FetchUserInfo(accessToken string, openid string) (*UserInfoResult, error) {
	url := fmt.Sprintf(userInfoUrl, accessToken, openid)
	body, err := utils.HTTPGet(url)
	if err != nil {
		return nil, err
	}
	var result UserInfoResult
	_ = json.Unmarshal(body, &result)
	if result.Openid == "" {
		global.GetLogger().Errorf("wechat app get user info is fail, openid:%s, error:%s", openid, string(body))
		return nil, errors.New("wechat app get user info is error")
	}
	return &result, nil
}

func (c *WeChatAppClient) CacheCodeAuth(code string, result *CodeAuthResult) {
	bytes, _ := json.Marshal(result)
	secs := time.Second * time.Duration(result.ExpiresIn)
	cache.RedisClient.Set(static.GetWxAppCodeAuthKey(code), string(bytes), secs)
}

func (c *WeChatAppClient) GetCodeAuth(code string) *CodeAuthResult {
	body := cache.RedisClient.Get(static.GetWxAppCodeAuthKey(code))
	if body == nil || cast.ToString(body) == "" {
		return nil
	}
	var ret CodeAuthResult
	json.Unmarshal([]byte(cast.ToString(body)), &ret)
	return &ret
}
