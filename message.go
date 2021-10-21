package wxcom

import (
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

func (m *Message) body() map[string]interface{} {

	if m.touser == "" && m.toparty == "" && m.totag == "" {
		panic("touser, toparty, totag cannot be empty at the same time")
	}

	body := map[string]interface{}{
		"touser":                   m.touser,
		"toparty":                  m.toparty,
		"totag":                    m.totag,
		"agentid":                  m.agentid,
		"enable_duplicate_check":   m.enableDuplicateCheck,
		"duplicate_check_interval": m.duplicateCheckInterval,
	}

	return body
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

	if response.Errcode == 42001 {
		m.wxcom.setAccessToken(m.wxcom.getAccessTokenFromServer().AccessToken)
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
func (m *Message) Text(content string) *Text {
	return &Text{
		message: m,
		content: content,
	}
}

// Image method creates image message.
func (m *Message) Image(mediaId string) *Image {
	return &Image{
		message: m,
		mediaId: mediaId,
	}
}

// Voice method creates voice message.
func (m *Message) Voice(mediaId string) *Voice {
	return &Voice{
		message: m,
		mediaId: mediaId,
	}
}

// Video method creates video message.
func (m *Message) Video(mediaId, title, description string) *Video {
	return &Video{
		message:     m,
		mediaId:     mediaId,
		title:       title,
		description: description,
	}
}

// File method creates file message.
func (m *Message) File(mediaId string) *File {
	return &File{
		message: m,
		mediaId: mediaId,
	}
}

// Textcard method creates textcard message.
func (m *Message) Textcard(title, description, url string) *Textcard {
	return &Textcard{
		message:     m,
		title:       title,
		description: description,
		url:         url,
	}
}

// Markdown method creates markdown message.
func (m *Message) Markdown(content string) *Markdown {
	return &Markdown{
		message: m,
		content: content,
	}
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Text
//__________________________________________________

// Text struct is used to compose txt message push from message client.
type Text struct {
	message       *Message
	content       string
	safe          int
	enableIdTrans int
}

// SetSafe method sets the message is confident.
func (t *Text) SetSafe(safe int) *Text {
	t.safe = safe
	return t
}

// SetEnableIdTrans method sets the message enable id translation.
func (t *Text) SetEnableIdTrans(enableIdTrans int) *Text {
	t.enableIdTrans = enableIdTrans
	return t
}

// Send method does send message.
func (t *Text) Send() (*MessageResponse, error) {
	body := t.message.body()
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
// Image
//__________________________________________________

// Image struct is used to compose image message push from message client.
type Image struct {
	message *Message
	mediaId string
	safe    int
}

// SetSafe method sets the message is confident.
func (i *Image) SetSafe(safe int) *Image {
	i.safe = safe
	return i
}

// Send method does send message.
func (i *Image) Send() (*MessageResponse, error) {
	body := i.message.body()
	body["msgtype"] = "image"
	body["image"] = map[string]string{"media_id": i.mediaId}

	resp, err := i.message.send(body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Voice
//__________________________________________________

// Voice struct is used to compose voice message push from message client.
type Voice struct {
	message *Message
	mediaId string
}

// Send method does send message.
func (v *Voice) Send() (*MessageResponse, error) {
	body := v.message.body()
	body["msgtype"] = "voice"
	body["voice"] = map[string]string{"media_id": v.mediaId}

	resp, err := v.message.send(body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Video
//__________________________________________________

// Video struct is used to compose video message push from message client.
type Video struct {
	message     *Message
	mediaId     string
	title       string
	description string
	safe        int
}

// SetSafe method sets the message is confident.
func (v *Video) SetSafe(safe int) *Video {
	v.safe = safe
	return v
}

// Send method does send message.
func (v *Video) Send() (*MessageResponse, error) {
	body := v.message.body()
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
// File
//__________________________________________________

// File struct is used to compose file message push from message client.
type File struct {
	message *Message
	mediaId string
	safe    int
}

// SetSafe method sets the message is confident.
func (f *File) SetSafe(safe int) *File {
	f.safe = safe
	return f
}

// Send method does send message.
func (f *File) Send() (*MessageResponse, error) {
	body := f.message.body()
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
// Textcard
//__________________________________________________

// Textcard struct is used to compose textcard message push from message client.
type Textcard struct {
	message       *Message
	title         string
	description   string
	url           string
	btntxt        string
	enableIdTrans int
}

// SetEnableIdTrans method sets the message enable id translation.
func (t *Textcard) SetEnableIdTrans(enableIdTrans int) *Textcard {
	t.enableIdTrans = enableIdTrans
	return t
}

// SetBtntxt method sets the textcard message btn txt.
func (t *Textcard) SetBtntxt(enableIdTrans int) *Textcard {
	t.enableIdTrans = enableIdTrans
	return t
}

// Send method does send message.
func (t *Textcard) Send() (*MessageResponse, error) {
	body := t.message.body()
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
// Markdown
//__________________________________________________

// Markdown struct is used to compose markdown message push from message client.
type Markdown struct {
	message *Message
	content string
}

// Send method does send message.
func (m *Markdown) Send() (*MessageResponse, error) {
	body := m.message.body()
	body["msgtype"] = "markdown"
	body["markdown"] = map[string]string{"content": m.content}

	resp, err := m.message.send(body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
