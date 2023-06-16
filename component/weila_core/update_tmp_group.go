package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type UpdateTmpGroupReq struct {
	GroupId int    `json:"group_id"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
}

func (s *WeilaService) UpdateTmpGroup(baseParam dto.BaseParam, req *UpdateTmpGroupReq) error {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/corp-group/change-tmp-group?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return err
	}
	global.GetLogger().InfoByTrace(&baseParam, "UpdateTmpGroup req:%+v, resp: %s", req, string(body))
	var resp CoreResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "UpdateTmpGroup error: %+v", resp)
		return resp.Error()
	}
	return nil
}
