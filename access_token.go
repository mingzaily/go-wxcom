package wxcom

import (
	"fmt"
	"time"
)

type respAccessToken struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (c *Client) getAccessTokenFromServer() *respAccessToken {
	var response *respAccessToken

	if c.corpid == "" && c.corpsecret == "" {
		panic("corpid and corpsecret cannot be empty")
	}

	_, err := c.Resty.R().
		SetQueryParam("corpid", c.corpid).
		SetQueryParam("corpsecret", c.corpsecret).
		SetResult(&response).
		Get("/cgi-bin/gettoken")
	if err != nil {
		panic(err)
	}

	if response.Errcode != 0 {
		panic(response.Errmsg)
	}

	return response
}

func (c *Client) GetAccessToken() string {
	var cacheKey = "access_token_" + fmt.Sprintf("%d", c.agentid)

	if value, found := c.cache.Get(cacheKey); found {
		return value.(string)
	}

	resp := c.getAccessTokenFromServer()
	c.cache.Set(cacheKey, resp.AccessToken, time.Duration(resp.ExpiresIn-60)*time.Second)

	return resp.AccessToken
}
