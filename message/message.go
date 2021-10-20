package message

import (
	wxcom "github.com/mingzaily/wxcom-sdk"
)

type Message struct {
	client *wxcom.Client
	url    string
}

type SendMessageRequest interface {
	Sendable() bool
	SetAgentid(agentid int)
}

func NewWithClient(client *wxcom.Client) *Message {
	return &Message{
		client: client,
		url:    "https://qyapi.weixin.qq.com/cgi-bin/message/send",
	}
}

func (m *Message) SendAppMessage(request SendMessageRequest) (*AppMessageResponse, error) {
	var response *AppMessageResponse

	if !request.Sendable() {
		panic("touser, toparty, totag cannot be empty at the same time")
	}

	request.SetAgentid(m.client.GetAgentid())

	_, err := m.client.C().R().
		SetQueryParam("access_token", m.client.GetAccessToken()).
		SetHeader("Content-Type", "application/json; charset=UTF-8").
		SetBody(request).
		SetResult(&response).
		Post(m.url)
	if err != nil {
		return nil, err
	}

	return response, nil
}
