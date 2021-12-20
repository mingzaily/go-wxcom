package wxcom

import (
	"fmt"
	"net/url"
)

// Oauth struct is used to get user info client.
type Oauth struct {
	wx   *Wxcom
	path string
}

// RespOauth struct holds response values of get user info.
type RespOauth struct {
	respCommon
	UserId         string `json:"UserId"`
	DeviceId       string `json:"DeviceId"`
	OpenId         string `json:"OpenId"`
	ExternalUserId string `json:"external_userid"`
}

// GenAuthorizationUrl method to construct web page authorization link
func (o *Oauth) GenAuthorizationUrl(redirectUri string) string {
	return fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize"+
		"?appid=%s"+
		"&redirect_uri=%s"+
		"&response_type=code"+
		"&scope=snsapi_base"+
		"&state="+
		"#wechat_redirect", o.wx.corpid, url.QueryEscape(redirectUri))
}

// GenAuthorizeScanCodeUrl method to build the scan code login authorization link
func (o *Oauth) GenAuthorizeScanCodeUrl(redirectUri string) string {
	return fmt.Sprintf("https://open.work.weixin.qq.com/wwopen/sso/qrConnect"+
		"?appid=%s"+
		"&agentid=%d"+
		"&redirect_uri=%s"+
		"&state=", o.wx.corpid, o.wx.agentid, url.QueryEscape(redirectUri))
}

func (o *Oauth) GetUserInfo(code string) (*RespOauth, error) {
	response := &RespOauth{}

	err := o.wx.sendWithRetry(o.path, map[string]string{"code": code}, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
