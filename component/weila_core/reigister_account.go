package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type RegisterAccountReq struct {
	Phone      string `json:"phone"`
	VerifyCode string `json:"verify_code"`
}

type RegisterAccountResp struct {
	CoreResp
	Data RegisterAccountData `json:"data"`
}

type RegisterAccountData struct {
	UserId     int    `json:"user_id"`
	UserNumber string `json:"user_number"`
	Password   string `json:"password"`
}

// RegisterAccount @Description: 注册账号
func (s *WeilaService) RegisterAccount(baseParam dto.BaseParam, req *RegisterAccountReq) (*RegisterAccountData, error) {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/user/regist-account?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	global.GetLogger().InfoByTrace(&baseParam, "RegisterAccount req:%+v, resp: %s", req, string(body))
	var resp RegisterAccountResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "RegisterAccount error: %+v", resp)
		return nil, resp.Error()
	}
	return &resp.Data, nil
}
