package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type CreateTmpGroupReq struct {
	OwnerId   int    `json:"owner_id"`
	Name      string `json:"name"`
	MemberIds []int  `json:"member_ids"`
}

type CreateTmpGroupResp struct {
	CoreResp
	Data struct {
		Group CreateTmpGroupData `json:"group"`
	} `json:"data"`
}

type CreateTmpGroupData struct {
	CreateCorpGroupData
}

func (s *WeilaService) CreateTmpGroup(baseParam dto.BaseParam, req *CreateTmpGroupReq) (*CreateTmpGroupData, error) {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/corp-group/create-tmp-group?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	global.GetLogger().InfoByTrace(&baseParam, "CreateTmpGroup req:%+v, resp: %s", req, string(body))
	var resp CreateTmpGroupResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "CreateTmpGroup error: %+v", resp)
		return nil, resp.Error()
	}
	return &resp.Data.Group, nil
}
