package utils

import (
	"fmt"
	"github.com/levigross/grequests"
	"strconv"
)

func StringToUint(s string) uint {
	result, _ := strconv.ParseUint(s, 10, 64)
	return uint(result)
}

func StringToInt(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}

type WXResponse struct {
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrMsg     string `json:"errmsg"`
	OpenID     string `json:"openid"`
	ErrCode    int    `json:"errcode"`
}

func GetSession(appID, appSecret, jsCode string) (*WXResponse, error) {
	apiURL := "https://api.weixin.qq.com/sns/jscode2session"

	// 设置请求参数
	params := map[string]string{
		"appid":      appID,
		"secret":     appSecret,
		"js_code":    jsCode,
		"grant_type": "authorization_code",
	}

	// 发送 GET 请求
	resp, err := grequests.Get(apiURL, &grequests.RequestOptions{Params: params})
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	// 解析响应 JSON
	var wxResp WXResponse
	if err := resp.JSON(&wxResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// 检查错误码
	if wxResp.ErrCode != 0 {
		return nil, fmt.Errorf("error code: %d, error message: %s", wxResp.ErrCode, wxResp.ErrMsg)
	}

	return &wxResp, nil
}
