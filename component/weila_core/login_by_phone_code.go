package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type LoginByPhoneCodeReq struct {
	Phone      string `json:"phone"`
	VerifyCode string `json:"verify_code"`
}

type LoginByPhoneCodeResp struct {
	CoreResp
	Data LoginByPhoneCodeData `json:"data"`
}

type LoginByPhoneCodeData struct {
	Password string `json:"password"`
}

// GetPasswordByPhoneCode @Description: 手机号验证码获取登录密码
func (s *WeilaService) GetPasswordByPhoneCode(baseParam dto.BaseParam, req *LoginByPhoneCodeReq) (*LoginByPhoneCodeData, error) {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/user/get-login-password?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	global.GetLogger().InfoByTrace(&baseParam, "GetPasswordByPhoneCode req:%+v, resp: %s", req, string(body))
	var resp LoginByPhoneCodeResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "GetPasswordByPhoneCode error: %+v", resp)
		return nil, resp.Error()
	}
	return &resp.Data, nil
}
