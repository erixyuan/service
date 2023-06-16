package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type CreateAccountReq struct {
	Password string `json:"password"`
}

type CreateAccountResp struct {
	CoreResp
	Data CreateAccountData `json:"data"`
}

type CreateAccountData struct {
	UserId     int    `json:"user_id"`
	UserNumber string `json:"user_number"`
}

// CreateAccount @Description:创建账号
func (s *WeilaService) CreateAccount(baseParam dto.BaseParam, req *CreateAccountReq) (*CreateAccountData, error) {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/user/create-account?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	global.GetLogger().InfoByTrace(&baseParam, "CreateAccount req:%+v, resp: %s", req, string(body))
	var resp CreateAccountResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "CreateAccount error: %+v", resp)
		return nil, resp.Error()
	}
	return &resp.Data, nil
}
