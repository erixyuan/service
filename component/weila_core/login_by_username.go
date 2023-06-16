package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type LoginByUsernameReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// LoginByUsername @Description:微喇号密码登录
func (s *WeilaService) LoginByUsername(baseParam dto.BaseParam, req *LoginByUsernameReq) (*LoginData, error) {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/user/login-by-username?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	global.GetLogger().InfoByTrace(&baseParam, "LoginByUsername req:%+v, resp: %s", req, string(body))
	var resp LoginCommonResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "LoginByUsername error: %+v", resp)
		return nil, resp.Error()
	}
	return &resp.Data, nil
}
