package wxcom

// Message Notification push client
type Message struct {
	wxcom *Wxcom
	url   string
}

type MessageRequest interface {
	sendable() bool
	setAgentid(agentid int)
}

// SendAppMessage Send application push messages.
func (m *Message) SendAppMessage(request MessageRequest) (*AppMessageResponse, error) {
	var response *AppMessageResponse

	if !request.sendable() {
		panic("touser, toparty, totag cannot be empty at the same time")
	}

	request.setAgentid(m.wxcom.agentid)

	_, err := m.wxcom.resty.R().
		SetQueryParam("access_token", m.wxcom.getAccessToken()).
		SetHeader("Content-Type", "application/json; charset=UTF-8").
		SetBody(request).
		SetResult(&response).
		Post(m.url)
	if err != nil {
		return nil, err
	}

	return response, nil
}
