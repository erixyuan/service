package lanpice_dto

type WeilaServiceInfo struct {
	ServiceId     int    `json:"service_id"`
	ServiceNumber string `json:"service_number"`
	Name          string `json:"name"`
	UserId        int    `json:"user_id"`
	Category      int    `json:"category"`
	Avatar        string `json:"avatar"`
	Intro         string `json:"intro"`
	AreaCode      string `json:"area_code"`
	ServiceUrl    string `json:"service_url"`
	Status        int    `json:"status"`
	Created       int    `json:"created"`
}
