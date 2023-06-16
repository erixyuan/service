package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type LoginByPhonePwdReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginCommonResp struct {
	CoreResp
	Data LoginData `json:"data"`
}

type LoginData struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	UserId       int    `json:"user_id"`
	UserNumber   string `json:"user_number"`
}

// LoginByPhonePwd @Description:手机号密码登录
func (s *WeilaService) LoginByPhonePwd(baseParam dto.BaseParam, req *LoginByPhonePwdReq) (*LoginData, error) {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/user/login-by-phone?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	global.GetLogger().InfoByTrace(&baseParam, "LoginByPhonePwd req:%+v, resp: %s", req, string(body))
	var resp LoginCommonResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "LoginByPhonePwd error: %+v", resp)
		return nil, resp.Error()
	}
	return &resp.Data, nil
}
