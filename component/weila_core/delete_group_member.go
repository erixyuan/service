package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type DeleteGroupMemberReq struct {
	GroupId   int   `json:"group_id"`
	MemberIds []int `json:"member_ids"`
}

// DeleteGroupMember @Description: 删除群成员
func (s *WeilaService) DeleteGroupMember(baseParam dto.BaseParam, req *DeleteGroupMemberReq) error {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/corp-group/delete-group-member?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return err
	}
	global.GetLogger().InfoByTrace(&baseParam, "DeleteGroupMember req:%+v, resp: %s", req, string(body))
	var resp CoreResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "DeleteGroupMember error: %+v", resp)
		return resp.Error()
	}
	return nil
}
