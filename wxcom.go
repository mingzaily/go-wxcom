package wxcom

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
	"time"
)

// Wxcom struct is used to create wxcom client.
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
	Resty      *resty.Client
}

type respAccessToken struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// New method creates a new Wxcom client.
func New(corpid, corpsecret string, agentid int) *Wxcom {
	return &Wxcom{
		corpid:     corpid,
		corpsecret: corpsecret,
		agentid:    agentid,
		cache:      cache.New(5*time.Minute, 10*time.Minute),
		Resty:      resty.New().SetBaseURL("https://qyapi.weixin.qq.com/"),
	}
}

func (w *Wxcom) getAccessTokenFromServer() *respAccessToken {
	var response *respAccessToken

	if w.corpid == "" && w.corpsecret == "" {
		panic("corpid and corpsecret cannot be empty")
	}

	_, err := w.Resty.R().
		SetQueryParam("corpid", w.corpid).
		SetQueryParam("corpsecret", w.corpsecret).
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

// isTokenInvalidErr method check whether the token has expired.
func (w *Wxcom) isTokenInvalidErr(errcode int) bool {
	if errcode == 42001 || errcode == 40014 {
		w.cache.Delete("access_token_" + fmt.Sprintf("%d", w.agentid))
		return true
	}
	return false
}

// GetAccessToken method get access token.
func (w *Wxcom) GetAccessToken() string {
	var cacheKey = "access_token_" + fmt.Sprintf("%d", w.agentid)

	if value, found := w.cache.Get(cacheKey); found {
		return value.(string)
	}

	resp := w.getAccessTokenFromServer()
	w.cache.Set(cacheKey, resp.AccessToken, time.Duration(resp.ExpiresIn-60)*time.Second)

	return resp.AccessToken
}

// GetAgentid method get agentid from client
func (w *Wxcom) GetAgentid() int {
	return w.agentid
}

// M method creates a new Message instance.
func (w *Wxcom) M() *Message {
	return &Message{
		wx:   w,
		path: "/cgi-bin/message/send",
	}
}

// NewMessage is an alias for method `M()`. Creates a new Message instance.
func (w *Wxcom) NewMessage() *Message {
	return w.M()
}
