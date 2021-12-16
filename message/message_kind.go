package message

type KindMessage interface {
	ToJson() string
	Send() (*RespMessage, error)
}

// text struct is used to compose text message push from message client.
type text struct {
	message *Message
	content string
}

// build method create the new Message client.
func (t *text) build() *Message {
	msg := t.message.clone()
	msg.msgType = "text"
	msg.content = t.content
	return msg
}

// ToJson method return text message string.
func (t *text) ToJson() string {
	return t.build().toJson()
}

// Send method does Send text message.
func (t *text) Send() (*RespMessage, error) {
	return t.build().send()
}

// image struct is used to compose image message push from message client.
type image struct {
	message *Message
	mediaId string
}

// build method create the new Message client.
func (i *image) build() *Message {
	msg := i.message.clone()
	msg.msgType = "image"
	msg.mediaId = i.mediaId
	return msg
}

// ToJson method return image message string.
func (i *image) ToJson() string {
	return i.build().toJson()
}

// Send method does send image message.
func (i *image) Send() (*RespMessage, error) {
	return i.build().send()
}

// voice struct is used to compose voice message push from message client.
type voice struct {
	message *Message
	mediaId string
}

// build method create the new Message client.
func (v *voice) build() *Message {
	msg := v.message.clone()
	msg.msgType = "voice"
	msg.mediaId = v.mediaId
	return msg
}

// ToJson method return voice message string.
func (v *voice) ToJson() string {
	return v.build().toJson()
}

// Send method does Send voice message.
func (v *voice) Send() (*RespMessage, error) {
	return v.build().send()
}

// video struct is used to compose video message push from message client.
type video struct {
	message     *Message
	mediaId     string
	title       string
	description string
}

// build method create the new Message client.
func (v *video) build() *Message {
	msg := v.message.clone()
	msg.msgType = "video"
	msg.mediaId = v.mediaId
	msg.title = v.title
	msg.description = v.description
	return msg
}

// SetTitle method sets the video message title.
func (v *video) SetTitle(title string) *video {
	v.title = title
	return v
}

// SetDescription method sets the video message description.
func (v *video) SetDescription(description string) *video {
	v.description = description
	return v
}

// ToJson method return video message string.
func (v *video) ToJson() string {
	return v.build().toJson()
}

// Send method does send video message.
func (v *video) Send() (*RespMessage, error) {
	return v.build().send()
}

// file struct is used to compose file message push from message client.
type file struct {
	message *Message
	mediaId string
}

// build method create the new Message client.
func (f *file) build() *Message {
	msg := f.message.clone()
	msg.mediaId = f.mediaId
	return msg
}

// ToJson method return file message string.
func (f *file) ToJson() string {
	return f.build().toJson()
}

// Send method does send file message.
func (f *file) Send() (*RespMessage, error) {
	return f.build().send()
}

// textcard struct is used to compose textcard message push from message client.wx_message
type textcard struct {
	message     *Message
	title       string
	description string
	url         string
	btnTxt      string
}

// build method create the new Message client.
func (t *textcard) build() *Message {
	msg := t.message.clone()
	msg.msgType = "textcard"
	msg.title = t.title
	msg.description = t.description
	msg.url = t.url
	msg.btnTxt = t.btnTxt
	return msg
}

// SetBtnTxt method sets the textcard message btn txt.
func (t *textcard) SetBtnTxt(btnTxt string) *textcard {
	t.btnTxt = btnTxt
	return t
}

// ToJson method return textcard message string.
func (t *textcard) ToJson() string {
	return t.build().toJson()
}

// Send method does send textcard message.
func (t *textcard) Send() (*RespMessage, error) {
	return t.build().send()
}

// markdown struct is used to compose markdown message push from message client.
type markdown struct {
	message *Message
	content string
}

// build method create the new Message client.
func (m *markdown) build() *Message {
	msg := m.message.clone()
	msg.msgType = "markdown"
	msg.content = m.content
	return msg
}

// ToJson method return markdown message string.
func (m *markdown) ToJson() string {
	return m.build().toJson()
}

// Send method does Send markdown message.
func (m *markdown) Send() (*RespMessage, error) {
	return m.build().send()
}
