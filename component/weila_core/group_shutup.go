package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type GroupShutupReq struct {
	GroupId int `json:"group_id"`
	Status  int `json:"status"`
}

// GroupShutup @Description:群禁言
func (s *WeilaService) GroupShutup(baseParam dto.BaseParam, req *GroupShutupReq) error {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/corp-group/group-shutup?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return err
	}
	global.GetLogger().InfoByTrace(&baseParam, "GroupShutup req:%+v, resp: %s", req, string(body))
	var resp CoreResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "GroupShutup error: %+v", resp)
		return resp.Error()
	}
	return nil
}
