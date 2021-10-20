package wx_com

import (
	"errors"
	"fmt"
	"time"
)

type accessTokenResponse struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (c *Client) getAccessTokenFromServer() (*accessTokenResponse, error) {
	var response *accessTokenResponse

	_, err := c.http.R().
		SetQueryParam("corpid", c.corpid).
		SetQueryParam("corpsecret", c.corpsecret).
		SetResult(&response).
		Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken")
	if err != nil {
		return nil, err
	}
	if response.Errcode != 0 {
		return nil, errors.New(response.Errmsg)
	}

	return response, nil
}

func (c *Client) GetAccessToken() string {
	var cacheKey = "access_token_" + fmt.Sprintf("%d", c.agentid)

	if value, found := c.cache.Get(cacheKey); found {
		return value.(string)
	}

	response, err := c.getAccessTokenFromServer()
	if err != nil {
		panic(err)
	}

	c.cache.Set(cacheKey, response.AccessToken, time.Duration(response.ExpiresIn-60)*time.Second)

	return response.AccessToken
}
