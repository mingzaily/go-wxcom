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

func (a *App) getAccessTokenFromServer() *respAccessToken {
	var response *respAccessToken

	if a.corpid == "" && a.corpsecret == "" {
		panic("corpid and corpsecret cannot be empty")
	}

	_, err := a.Resty.R().
		SetQueryParam("corpid", a.corpid).
		SetQueryParam("corpsecret", a.corpsecret).
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

func (a *App) GetAccessToken() string {
	var cacheKey = "access_token_" + fmt.Sprintf("%d", a.agentid)

	if value, found := a.cache.Get(cacheKey); found {
		return value.(string)
	}

	resp := a.getAccessTokenFromServer()
	a.cache.Set(cacheKey, resp.AccessToken, time.Duration(resp.ExpiresIn-60)*time.Second)

	return resp.AccessToken
}
