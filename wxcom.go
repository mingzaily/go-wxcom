package wxcom

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
	"time"
)

// Wxcom struct is used to create Wxcom client.
//
// The cache uses patrickmn/go-cache.
// You can refer to related documents(https://github.com/patrickmn/go-cache) if necessary.
//
// The resty uses go-resty/resty/v2.
// You can refer to related documents(https://github.com/go-resty/resty) if necessary.
type Wxcom struct {
	corpid     string
	corpsecret string
	agentid    int
	cache      *cache.Cache
	resty      *resty.Client
}

// New method creates a new Wxcom client.
func New(corpid, corpsecret string, agentid int) *Wxcom {
	return &Wxcom{
		corpid:     corpid,
		corpsecret: corpsecret,
		agentid:    agentid,
		cache:      cache.New(5*time.Minute, 10*time.Minute),
		resty:      resty.New(),
	}
}

// SetRestyDebug method enables the debug mode on Resty client. Client logs details of every request and response.
// For `Request` it logs information such as HTTP verb, Relative URL path, Host, Headers, Body if it has one.
// For `Response` it logs information such as Status, Response Time, Headers, Body if it has one.
func (w *Wxcom) SetRestyDebug(d bool) *Wxcom {
	w.resty.SetDebug(d)
	return w
}

type accessTokenResponse struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (w *Wxcom) getAccessTokenFromServer() *accessTokenResponse {
	var response *accessTokenResponse

	if w.corpid == "" && w.corpsecret == "" {
		panic("corpid and corpsecret cannot be empty")
	}

	_, err := w.resty.R().
		SetQueryParam("corpid", w.corpid).
		SetQueryParam("corpsecret", w.corpsecret).
		SetResult(&response).
		Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken")
	if err != nil {
		panic(err)
	}
	if response.Errcode != 0 {
		panic(response.Errmsg)
	}

	return response
}

func (w *Wxcom) getAccessToken() string {
	var cacheKey = "access_token_" + fmt.Sprintf("%d", w.agentid)

	if value, found := w.cache.Get(cacheKey); found {
		return value.(string)
	}

	resp := w.getAccessTokenFromServer()

	w.cache.Set(cacheKey, resp.AccessToken, time.Duration(resp.ExpiresIn-60)*time.Second)

	return resp.AccessToken
}

// Message method creates a new message instance, its used for send message to user form App.
func (w *Wxcom) Message() *Message {
	return &Message{
		wxcom:   w,
		url:     "https://qyapi.weixin.qq.com/cgi-bin/message/send",
		agentid: w.agentid,
	}
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Common method
//__________________________________________________

// IsTokenInvalidErr method check whether the token has expired.
func IsTokenInvalidErr(errcode int, w *Wxcom) bool {
	if errcode == 42001 || errcode == 40014 {
		w.cache.Delete("access_token_" + fmt.Sprintf("%d", w.agentid))
		return true
	}
	return false
}
