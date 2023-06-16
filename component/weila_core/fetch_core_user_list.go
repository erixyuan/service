package lanpice_core

import (
	"encoding/json"
	"fmt"
	"service/dto"
	"service/global"
	"service/utils"
)

type CoreUserDTO struct {
	Id      int    `json:"id"`
	Number  string `json:"number"`
	Nick    string `json:"nick"`
	Avatar  string `json:"avatar"`
	Sex     int    `json:"sex"`
	Created string `json:"created"`
	Phone   string `json:"phone"`
}

type FetchCoreUserInfoResp struct {
	CoreResp
	Data struct {
		Users []*CoreUserDTO `json:"users"`
	}
}

type FetchCoreUserInfoRq struct {
	WeilaNoSets  *[]string `json:"user_numbers"`
	WeilaUserIds *[]int    `json:"user_ids"`
}

func (s *WeilaService) FetchCoreUserInfo(baseParam dto.BaseParam, lanpiceNo string) (*CoreUserDTO, error) {
	userIds := []string{lanpiceNo}
	info, err := s.FetchCoreUserListInfo(baseParam, userIds)
	if err != nil {
		return nil, err
	}
	return info[0], nil
}

// 通用获取组织成员信息接口
func (s *WeilaService) FetchCoreUserListInfo(baseParam dto.BaseParam, lanpiceNoSets []string) ([]*CoreUserDTO, error) {
	req := FetchCoreUserInfoRq{
		WeilaNoSets: &lanpiceNoSets,
	}
	return s.GetUserInfoList(baseParam, &req)
}

// 通用获取组织成员信息接口
func (s *WeilaService) FetchCoreUserInfoByUserId(baseParam dto.BaseParam, lanpiceUserId int) (*CoreUserDTO, error) {
	userIds := []int{lanpiceUserId}
	info, err := s.FetchCoreUserListInfoByUserIds(baseParam, userIds)
	if err != nil {
		return nil, err
	}
	return info[0], nil
}

// 通用获取组织成员信息接口
func (s *WeilaService) FetchCoreUserListInfoByUserIds(baseParam dto.BaseParam, lanpiceUserIds []int) ([]*CoreUserDTO, error) {
	req := FetchCoreUserInfoRq{
		WeilaUserIds: &lanpiceUserIds,
	}
	return s.GetUserInfoList(baseParam, &req)
}

// 通用获取组织成员信息接口
func (s *WeilaService) GetUserInfoList(baseParam dto.BaseParam, req *FetchCoreUserInfoRq) ([]*CoreUserDTO, error) {
	global.GetLogger().InfoByTrace(&baseParam, "开始获取user_info_list: %+v", req)
	accessToken := s.FetchAccessToken(baseParam)
	url := fmt.Sprintf("%s/v1/third/server/user/get-user-info?appid=%s&token=%s", s.HostDomain, s.AppId, accessToken)
	body, err := utils.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	global.GetLogger().InfoByTrace(&baseParam, "FetchCoreUserListInfo req:%+v, resp: %s", req, string(body))
	var result FetchCoreUserInfoResp
	_ = json.Unmarshal(body, &result)
	if result.Error() != nil {
		global.GetLogger().ErrorByTrace(&baseParam, "FetchCoreUserListInfo error: %+v", result)
		return nil, result.Error()
	}
	return result.Data.Users, nil
}
