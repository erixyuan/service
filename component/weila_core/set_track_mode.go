package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type SetTrackModeReq struct {
	UserNumber string `json:"user_number"`
	Interval   int    `json:"interval"`
}

// SetTrackMode @Description:设置轨迹上报频率
func (s *WeilaService) SetTrackMode(baseParam dto.BaseParam, req *SetTrackModeReq) error {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/user/set-track-mode?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return err
	}
	global.GetLogger().InfoByTrace(&baseParam, "SetTrackMode req:%+v, resp: %s", req, string(body))
	var resp CoreResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "SetTrackMode error: %+v", resp)
		return resp.Error()
	}
	return nil
}
