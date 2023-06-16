package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type AddGroupMemberReq struct {
	GroupId   int   `json:"group_id"`
	MemberIds []int `json:"member_ids"`
}

// AddGroupMember @Description: 添加群成员
func (s *WeilaService) AddGroupMember(baseParam dto.BaseParam, req *AddGroupMemberReq) error {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/corp-group/add-group-member?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return err
	}
	global.GetLogger().InfoByTrace(&baseParam, "AddGroupMember req:%+v, resp: %s", req, string(body))
	var resp CoreResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "AddGroupMember error: %+v", resp)
		return resp.Error()
	}
	return nil
}
