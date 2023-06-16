package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type UpdateCorpGroupReq struct {
	GroupId int    `json:"group_id"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
}

func (s *WeilaService) UpdateCorpGroup(baseParam dto.BaseParam, req *UpdateCorpGroupReq) error {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/corp-group/change-normal-group?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return err
	}
	global.GetLogger().InfoByTrace(&baseParam, "UpdateCorpGroup req:%+v, resp: %s", req, string(body))
	var resp CoreResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "UpdateCorpGroup error: %+v", resp)
		return resp.Error()
	}
	return nil
}
