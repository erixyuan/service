package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type GetGroupMembersReq struct {
}

type GetGroupMembersResp struct {
	CoreResp
	Data GetGroupMembersData `json:"data"`
}

type GetGroupMembersData struct {
}

// GetGroupMembers @Description: 获取群组成员
func (s *WeilaService) GetGroupMembers(baseParam dto.BaseParam, req *GetGroupMembersReq) (*GetGroupMembersData, error) {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/group/get-group-members?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	global.GetLogger().InfoByTrace(&baseParam, "GetGroupMembers req:%+v, resp: %s", req, string(body))
	var resp GetGroupMembersResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "GetGroupMembers error: %+v", resp)
		return nil, resp.Error()
	}
	return &resp.Data, nil
}
