package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type DeleteTmpGroupReq struct {
	GroupId int `json:"group_id"`
}

func (s *WeilaService) DeleteTmpGroup(baseParam dto.BaseParam, req *DeleteTmpGroupReq) error {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/corp-group/delete-tmp-group?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return err
	}
	global.GetLogger().InfoByTrace(&baseParam, "DeleteTmpGroup req:%+v, resp: %s", req, string(body))
	var resp CoreResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "DeleteTmpGroup error: %+v", resp)
		return resp.Error()
	}
	return nil
}
