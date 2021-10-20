package wxcom

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
	"sync"
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
	mutex      *sync.RWMutex
}

// New method creates a new Wxcom client.
func New(corpid, corpsecret string, agentid int) *Wxcom {
	return &Wxcom{
		corpid:     corpid,
		corpsecret: corpsecret,
		agentid:    agentid,
		cache:      cache.New(5*time.Minute, 10*time.Minute),
		resty:      resty.New(),
		mutex:      &sync.RWMutex{},
	}
}

func (w *Wxcom) SetRequestDebug(d bool) *Wxcom {
	w.resty.SetDebug(d)
	return w
}

type accessTokenResponse struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (w *Wxcom) getAccessToken() string {
	var response *accessTokenResponse
	var cacheKey = "access_token_" + fmt.Sprintf("%d", w.agentid)

	if value, found := w.cache.Get(cacheKey); found {
		return value.(string)
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

	w.cache.Set(cacheKey, response.AccessToken, time.Duration(response.ExpiresIn-60)*time.Second)

	return response.AccessToken
}

// M create notice push client with sdk client
func (w *Wxcom) M() *Message {
	return &Message{
		wxcom: w,
		url:   "https://qyapi.weixin.qq.com/cgi-bin/message/send",
	}
}
