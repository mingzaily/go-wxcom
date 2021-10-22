package wxcom

import (
	"errors"
	"strings"
)

// Message struct is used to compose and fire individual message push from wxcom client.
type Message struct {
	wxcom                  *Wxcom
	url                    string
	touser                 string
	toparty                string
	totag                  string
	agentid                int
	enableDuplicateCheck   int
	duplicateCheckInterval int
}

// MessageResponse struct holds response values of send message.
type MessageResponse struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
	Msgid        string `json:"msgid"`
	ResponseCode string `json:"response_code"`
}

func (m *Message) body() (map[string]interface{}, error) {

	if m.touser == "" && m.toparty == "" && m.totag == "" {
		return nil, errors.New("touser, toparty, totag cannot be empty at the same time")
	}

	body := map[string]interface{}{
		"touser":                   m.touser,
		"toparty":                  m.toparty,
		"totag":                    m.totag,
		"agentid":                  m.agentid,
		"enable_duplicate_check":   m.enableDuplicateCheck,
		"duplicate_check_interval": m.duplicateCheckInterval,
	}

	return body, nil
}

func (m *Message) send(body map[string]interface{}) (*MessageResponse, error) {
	var response *MessageResponse

	_, err := m.wxcom.resty.R().
		SetQueryParam("access_token", m.wxcom.getAccessToken()).
		SetHeader("Content-Type", "application/json; charset=UTF-8").
		SetBody(body).
		SetResult(&response).
		Post(m.url)
	if err != nil {
		return nil, err
	}

	if IsTokenInvalidErr(response.Errcode, m.wxcom) {
		return m.send(body)
	}

	return response, nil
}

// SetTouser method sets to user to in the current message.
func (m *Message) SetTouser(userList []string) *Message {
	m.touser = strings.Join(userList, "|")
	return m
}

// SetToparty method sets to party to in the current message.
func (m *Message) SetToparty(partyList []string) *Message {
	m.toparty = strings.Join(partyList, "|")
	return m
}

// SetTotag method sets to tag to in the current message.
func (m *Message) SetTotag(tagList []string) *Message {
	m.totag = strings.Join(tagList, "|")
	return m
}

// SetDuplicateCheck method enables the duplicate check.
// Param example:
// 0, 0 duplicate check is not enabled.
// 1, 1800 duplicate check within 1800 seconds.
func (m *Message) SetDuplicateCheck(enableDuplicateCheck, duplicateCheckInterval int) *Message {
	m.enableDuplicateCheck = enableDuplicateCheck
	if m.enableDuplicateCheck == 1 {
		m.duplicateCheckInterval = duplicateCheckInterval
	}
	return m
}

// Text method creates text message.
func (m *Message) Text(content string) *text {
	return &text{
		message: m,
		content: content,
	}
}

// Image method creates image message.
func (m *Message) Image(mediaId string) *image {
	return &image{
		message: m,
		mediaId: mediaId,
	}
}

// Voice method creates voice message.
func (m *Message) Voice(mediaId string) *voice {
	return &voice{
		message: m,
		mediaId: mediaId,
	}
}

// Video method creates video message.
func (m *Message) Video(mediaId, title, description string) *video {
	return &video{
		message:     m,
		mediaId:     mediaId,
		title:       title,
		description: description,
	}
}

// File method creates file message.
func (m *Message) File(mediaId string) *file {
	return &file{
		message: m,
		mediaId: mediaId,
	}
}

// Textcard method creates textcard message.
func (m *Message) Textcard(title, description, url string) *textcard {
	return &textcard{
		message:     m,
		title:       title,
		description: description,
		url:         url,
	}
}

// Markdown method creates markdown message.
func (m *Message) Markdown(content string) *markdown {
	return &markdown{
		message: m,
		content: content,
	}
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// text
//__________________________________________________

// text struct is used to compose txt message push from message client.
type text struct {
	message       *Message
	content       string
	safe          int
	enableIdTrans int
}

// SetSafe method sets the message is confident.
func (t *text) SetSafe(safe int) *text {
	t.safe = safe
	return t
}

// SetEnableIdTrans method sets the message enable id translation.
func (t *text) SetEnableIdTrans(enableIdTrans int) *text {
	t.enableIdTrans = enableIdTrans
	return t
}

// Send method does send message.
func (t *text) Send() (*MessageResponse, error) {
	body, err := t.message.body()
	if err != nil {
		return nil, err
	}
	body["msgtype"] = "text"
	body["text"] = map[string]string{"content": t.content}
	body["safe"] = t.safe
	body["enable_id_trans"] = t.enableIdTrans

	resp, err := t.message.send(body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// image
//__________________________________________________

// image struct is used to compose image message push from message client.
type image struct {
	message *Message
	mediaId string
	safe    int
}

// SetSafe method sets the message is confident.
func (i *image) SetSafe(safe int) *image {
	i.safe = safe
	return i
}

// Send method does send message.
func (i *image) Send() (*MessageResponse, error) {
	body, err := i.message.body()
	if err != nil {
		return nil, err
	}
	body["msgtype"] = "image"
	body["image"] = map[string]string{"media_id": i.mediaId}
	body["safe"] = i.safe

	resp, err := i.message.send(body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// voice
//__________________________________________________

// voice struct is used to compose voice message push from message client.
type voice struct {
	message *Message
	mediaId string
}

// Send method does send message.
func (v *voice) Send() (*MessageResponse, error) {
	body, err := v.message.body()
	if err != nil {
		return nil, err
	}
	body["msgtype"] = "voice"
	body["voice"] = map[string]string{"media_id": v.mediaId}

	resp, err := v.message.send(body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// video
//__________________________________________________

// video struct is used to compose video message push from message client.
type video struct {
	message     *Message
	mediaId     string
	title       string
	description string
	safe        int
}

// SetSafe method sets the message is confident.
func (v *video) SetSafe(safe int) *video {
	v.safe = safe
	return v
}

// Send method does send message.
func (v *video) Send() (*MessageResponse, error) {
	body, err := v.message.body()
	if err != nil {
		return nil, err
	}
	body["msgtype"] = "video"
	body["video"] = map[string]string{"media_id": v.mediaId, "title": v.title, "description": v.description}
	body["safe"] = v.safe

	resp, err := v.message.send(body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// file
//__________________________________________________

// file struct is used to compose file message push from message client.
type file struct {
	message *Message
	mediaId string
	safe    int
}

// SetSafe method sets the message is confident.
func (f *file) SetSafe(safe int) *file {
	f.safe = safe
	return f
}

// Send method does send message.
func (f *file) Send() (*MessageResponse, error) {
	body, err := f.message.body()
	if err != nil {
		return nil, err
	}
	body["msgtype"] = "file"
	body["file"] = map[string]string{"media_id": f.mediaId}
	body["safe"] = f.safe

	resp, err := f.message.send(body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// textcard
//__________________________________________________

// textcard struct is used to compose textcard message push from message client.
type textcard struct {
	message       *Message
	title         string
	description   string
	url           string
	btntxt        string
	enableIdTrans int
}

// SetEnableIdTrans method sets the message enable id translation.
func (t *textcard) SetEnableIdTrans(enableIdTrans int) *textcard {
	t.enableIdTrans = enableIdTrans
	return t
}

// SetBtntxt method sets the textcard message btn txt.
func (t *textcard) SetBtntxt(btntxt string) *textcard {
	t.btntxt = btntxt
	return t
}

// Send method does send message.
func (t *textcard) Send() (*MessageResponse, error) {
	body, err := t.message.body()
	if err != nil {
		return nil, err
	}
	body["msgtype"] = "textcard"
	body["textcard"] = map[string]string{"title": t.title, "description": t.description, "url": t.url, "btntxt": t.btntxt}
	body["enable_id_trans"] = t.enableIdTrans

	resp, err := t.message.send(body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// markdown
//__________________________________________________

// markdown struct is used to compose markdown message push from message client.
type markdown struct {
	message *Message
	content string
}

// Send method does send message.
func (m *markdown) Send() (*MessageResponse, error) {
	body, err := m.message.body()
	if err != nil {
		return nil, err
	}
	body["msgtype"] = "markdown"
	body["markdown"] = map[string]string{"content": m.content}

	resp, err := m.message.send(body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
