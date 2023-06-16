package lanpice_core

import (
	"errors"
)

type CoreResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (r CoreResp) Error() error {
	if r.ErrCode == 0 {
		return nil
	}
	return errors.New(r.ErrMsg)
}
