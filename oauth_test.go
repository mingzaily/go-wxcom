package wxcom_test

import (
	"github.com/mingzaily/go-wxcom"
	"testing"
)

func TestOauth_GenAuthorizationUrl(t *testing.T) {
	oauth := wx.O()

	assertEqual(t,
		oauth.GenAuthorizationUrl("https://test.com/query"),
		"https://open.weixin.qq.com/connect/oauth2/authorize?appid=123&redirect_uri=https%3A%2F%2Ftest.com%2Fquery&response_type=code&scope=snsapi_base&state=#wechat_redirect")
}

func TestOauth_GenAuthorizeScanCodeUrl(t *testing.T) {
	oauth := wx.NewOauth()

	assertEqual(t,
		oauth.GenAuthorizeScanCodeUrl("https://test.com/query"),
		"https://open.work.weixin.qq.com/wwopen/sso/qrConnect?appid=123&agentid=1&redirect_uri=https%3A%2F%2Ftest.com%2Fquery&state=")
}

func TestOauth_GetUserInfo(t *testing.T) {
	ts := createTestServer(t)
	defer ts.Close()

	tempWx := wxcom.New("123", "321", 123)
	tempWx.Resty.SetBaseURL(ts.URL)

	resp, err := tempWx.O().GetUserInfo("code")

	assertEqual(t, err, nil)
	assertEqual(t, resp.Errcode, 0)
	assertEqual(t, resp.Errmsg, "ok")
	assertEqual(t, resp.UserId, "test_user")
	assertEqual(t, resp.DeviceId, "device")
}
