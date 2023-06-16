package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type DeleteCorpGroupReq struct {
	GroupId int `json:"group_id"`
}

func (s *WeilaService) DeleteCorpGroup(baseParam dto.BaseParam, req *DeleteCorpGroupReq) error {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/corp-group/delete-normal-group?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return err
	}
	global.GetLogger().InfoByTrace(&baseParam, "DeleteCorpGroup req:%+v, resp: %s", req, string(body))
	var resp CoreResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "DeleteCorpGroup error: %+v", resp)
		return resp.Error()
	}
	return nil
}
