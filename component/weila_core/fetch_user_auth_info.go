package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type FetchUserAuthInfoReq struct {
	AccessToken string `json:"access_token"`
}

type FetchUserAuthInfoResp struct {
	CoreResp
	Data struct {
		User *CoreUserDTO `json:"user"`
	} `json:"data"`
}

func (s *WeilaService) FetchUserAuthInfo(baseParam dto.BaseParam, req *FetchUserAuthInfoReq) (*CoreUserDTO, error) {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/oauth?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	var resp FetchUserAuthInfoResp
	_ = json.Unmarshal(body, &resp)
	global.GetLogger().InfoByTrace(&baseParam, "FetchUserAuthInfo req:%+v, resp: %s", req, string(body))
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "FetchUserAuthInfo error: %+v", resp)
		return nil, resp.Error()
	}
	return resp.Data.User, nil
}
