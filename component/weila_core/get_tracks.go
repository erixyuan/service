package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type GetTracksReq struct {
	UserNumber int    `json:"user_number"`
	Date       string `json:"date"`
}

type GetTracksResp struct {
	CoreResp
	Data struct {
		Tracks []*Track `json:"tracks"`
	} `json:"data"`
}

type Track struct {
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Created      string `json:"created"`
	LocationType string `json:"location_type"`
}

// GetTracks @Description: 获取轨迹
func (s *WeilaService) GetTracks(baseParam dto.BaseParam, req *GetTracksReq) ([]*Track, error) {
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/user/get-tracks?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	global.GetLogger().InfoByTrace(&baseParam, "GetTracks req:%+v, resp: %s", req, string(body))
	var resp GetTracksResp
	_ = json.Unmarshal(body, &resp)
	if resp.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "GetTracks error: %+v", resp)
		return nil, resp.Error()
	}
	return resp.Data.Tracks, nil
}
