package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type SendCodeReq struct {
	Phone string `json:"phone"`
}

// SendCode @Description: 发送短信验证码
func (s *WeilaService) SendCode(baseParam dto.BaseParam, req *SendCodeReq) error {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/user/send-verify-code?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return err
	}
	global.GetLogger().InfoByTrace(&baseParam, "SendCode req:%+v, resp: %s", req, string(body))
	var resp CoreResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "SendCode error: %+v", resp)
		return resp.Error()
	}
	return nil
}
