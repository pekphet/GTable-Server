package common

type BaseResp struct {
	ErrCode int `json:"errCode"`
	ErrMsg string `json:"errMsg"`
}