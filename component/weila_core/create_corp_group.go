package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type CreateCorpGroupReq struct {
	OwnerId int    `json:"owner_id"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar,omitempty"`
}

type CreateCorpGroupResp struct {
	CoreResp
	Data struct {
		Group CreateCorpGroupData `json:"group"`
	} `json:"data"`
}

type CreateCorpGroupData struct {
	Id        int    `json:"id"`
	OwnerId   int    `json:"owner_id"`
	Name      string `json:"name"`
	Number    string `json:"number"`
	Avatar    string `json:"avatar"`
	BurstMode int    `json:"burst_mode"`
	Shutup    int    `json:"shutup"`
	Type      int    `json:"type"`
}

// CreateCorpGroup @Description:创建企业群组
func (s *WeilaService) CreateCorpGroup(baseParam dto.BaseParam, req *CreateCorpGroupReq) (*CreateCorpGroupData, error) {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/corp-group/create-normal-group?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	global.GetLogger().InfoByTrace(&baseParam, "CreateCorpGroup req:%+v, resp: %s", req, string(body))
	var resp CreateCorpGroupResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "CreateCorpGroup error: %+v", resp)
		return nil, resp.Error()
	}
	return &resp.Data.Group, nil
}
