package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type ChangeGroupBurstModeReq struct {
	GroupId int `json:"group_id"`
	Mode    int `json:"mode"`
}

// ChangeGroupBurstMode @Description: 修改群对讲模式
func (s *WeilaService) ChangeGroupBurstMode(baseParam dto.BaseParam, req *ChangeGroupBurstModeReq) error {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/corp-group/change-group-burst-mode?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return err
	}
	global.GetLogger().InfoByTrace(&baseParam, "ChangeGroupBurstMode req:%+v, resp: %s", req, string(body))
	var resp CoreResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "ChangeGroupBurstMode error: %+v", resp)
		return resp.Error()
	}
	return nil
}
