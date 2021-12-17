package wxcom_test

import (
	"github.com/mingzaily/go-wxcom"
	"testing"
)

var msg = wxcom.New("123", "321", 123).M()

func TestMessage_NotToPeople(t *testing.T) {
	m := msg.Clone().Text("测试TEXT")

	_, err := m.Send()
	assertEqual(t,
		err.Error(),
		"toUser, toParty, toTag cannot be empty at the same time")

	assertEqual(t,
		m.ToJson(),
		"")
}

func TestMessage_Send(t *testing.T) {
	ts := createTestServer(t)
	defer ts.Close()

	tempWx := wxcom.New("123", "321", 123)
	tempWx.Resty.SetBaseURL(ts.URL)

	resp, err := tempWx.NewMessage().ToUser([]string{"test"}).Text("测试TEXT").Send()

	assertEqual(t, err, nil)
	assertEqual(t, resp.Errcode, 0)
	assertEqual(t, resp.Errmsg, "ok")
	assertEqual(t, resp.Msgid, "msgid")
}

func TestMessage_ToUser(t *testing.T) {
	m := msg.Clone().ToUser([]string{"user"}).Text("测试TEXT")

	assertEqual(t, m.ToJson(),
		"{\"agentid\":123,\"enable_id_trans\":0,\"msgtype\":\"text\",\"safe\":0,\"text\":{\"content\":\"测试TEXT\"},\"touser\":\"user\"}")
}

func TestMessage_ToParty(t *testing.T) {
	m := msg.Clone().ToParty([]string{"party"}).Text("测试TEXT")

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"enable_id_trans\":0,\"msgtype\":\"text\",\"safe\":0,\"text\":{\"content\":\"测试TEXT\"},\"toparty\":\"party\"}")
}

func TestMessage_ToTag(t *testing.T) {
	m := msg.Clone().ToTag([]string{"tag"}).Text("测试TEXT")

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"enable_id_trans\":0,\"msgtype\":\"text\",\"safe\":0,\"text\":{\"content\":\"测试TEXT\"},\"totag\":\"tag\"}")
}

func TestMessage_DuplicateCheck(t *testing.T) {
	m := msg.Clone().ToUser([]string{"user"}).DuplicateCheck(0, 1800).Text("测试TEXT")

	assertEqual(t, m.ToJson(),
		"{\"agentid\":123,\"enable_id_trans\":0,\"msgtype\":\"text\",\"safe\":0,\"text\":{\"content\":\"测试TEXT\"},\"touser\":\"user\"}")

	m = msg.Clone().ToUser([]string{"user"}).DuplicateCheck(1, 1800).Text("测试TEXT")

	assertEqual(t, m.ToJson(),
		"{\"agentid\":123,\"duplicate_check_interval\":1800,\"enable_duplicate_check\":1,\"enable_id_trans\":0,\"msgtype\":\"text\",\"safe\":0,\"text\":{\"content\":\"测试TEXT\"},\"touser\":\"user\"}")
}

func TestMessage_Text(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).Text("测试TEXT")

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"enable_id_trans\":0,\"msgtype\":\"text\",\"safe\":0,\"text\":{\"content\":\"测试TEXT\"},\"touser\":\"test\"}")
}

func TestMessage_Text_SetSafe(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).Text("测试TEXT").SetSafe(1)

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"enable_id_trans\":0,\"msgtype\":\"text\",\"safe\":1,\"text\":{\"content\":\"测试TEXT\"},\"touser\":\"test\"}")
}

func TestMessage_Text_SetEnableIdTrans(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).Text("测试TEXT").SetEnableIdTrans(1)

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"enable_id_trans\":1,\"msgtype\":\"text\",\"safe\":0,\"text\":{\"content\":\"测试TEXT\"},\"touser\":\"test\"}")
}

func TestMessage_Clone(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).Image("cac8d983-5d58-4603-9671-adb3edf2edf3")

	assertNotEqual(t, m, msg)
}

func TestMessage_Image(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).Image("cac8d983-5d58-4603-9671-adb3edf2edf3")

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"image\":{\"media_id\":\"cac8d983-5d58-4603-9671-adb3edf2edf3\"},\"msgtype\":\"image\",\"safe\":0,\"touser\":\"test\"}")
}

func TestMessage_Image_SetSafe(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).Image("cac8d983-5d58-4603-9671-adb3edf2edf3").SetSafe(1)

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"image\":{\"media_id\":\"cac8d983-5d58-4603-9671-adb3edf2edf3\"},\"msgtype\":\"image\",\"safe\":1,\"touser\":\"test\"}")
}

func TestMessage_Voice(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).Voice("cac8d983-5d58-4603-9671-adb3edf2edf3")

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"msgtype\":\"voice\",\"touser\":\"test\",\"voice\":{\"media_id\":\"cac8d983-5d58-4603-9671-adb3edf2edf3\"}}")
}

func TestMessage_Video(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).Video("cac8d983-5d58-4603-9671-adb3edf2edf3")

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"msgtype\":\"video\",\"safe\":0,\"touser\":\"test\",\"video\":{\"description\":\"\",\"media_id\":\"cac8d983-5d58-4603-9671-adb3edf2edf3\",\"title\":\"\"}}")
}

func TestMessage_Video_SetSafe(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).Video("cac8d983-5d58-4603-9671-adb3edf2edf3").SetSafe(1)

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"msgtype\":\"video\",\"safe\":1,\"touser\":\"test\",\"video\":{\"description\":\"\",\"media_id\":\"cac8d983-5d58-4603-9671-adb3edf2edf3\",\"title\":\"\"}}")
}

func TestMessage_Video_SetTitle(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).Video("cac8d983-5d58-4603-9671-adb3edf2edf3").SetTitle("标题")

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"msgtype\":\"video\",\"safe\":0,\"touser\":\"test\",\"video\":{\"description\":\"\",\"media_id\":\"cac8d983-5d58-4603-9671-adb3edf2edf3\",\"title\":\"标题\"}}")
}

func TestMessage_Video_SetDescription(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).Video("cac8d983-5d58-4603-9671-adb3edf2edf3").SetDescription("描述")

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"msgtype\":\"video\",\"safe\":0,\"touser\":\"test\",\"video\":{\"description\":\"描述\",\"media_id\":\"cac8d983-5d58-4603-9671-adb3edf2edf3\",\"title\":\"\"}}")
}

func TestMessage_File(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).File("cac8d983-5d58-4603-9671-adb3edf2edf3")

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"file\":{\"media_id\":\"cac8d983-5d58-4603-9671-adb3edf2edf3\"},\"msgtype\":\"file\",\"safe\":0,\"touser\":\"test\"}")
}

func TestMessage_File_SetSafe(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).File("cac8d983-5d58-4603-9671-adb3edf2edf3").SetSafe(1)

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"file\":{\"media_id\":\"cac8d983-5d58-4603-9671-adb3edf2edf3\"},\"msgtype\":\"file\",\"safe\":1,\"touser\":\"test\"}")
}

func TestMessage_Textcard(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).Textcard("标题", "描述", "url")

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"enable_id_trans\":0,\"msgtype\":\"textcard\",\"textcard\":{\"btntxt\":\"\",\"description\":\"描述\",\"title\":\"标题\",\"url\":\"url\"},\"touser\":\"test\"}")
}

func TestMessage_Textcard_SetBtnTxt(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).Textcard("标题", "描述", "url").SetBtnTxt("按钮")

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"enable_id_trans\":0,\"msgtype\":\"textcard\",\"textcard\":{\"btntxt\":\"按钮\",\"description\":\"描述\",\"title\":\"标题\",\"url\":\"url\"},\"touser\":\"test\"}")
}

func TestMessage_Textcard_SetEnableIdTrans(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).Textcard("标题", "描述", "url").SetEnableIdTrans(1)

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"enable_id_trans\":1,\"msgtype\":\"textcard\",\"textcard\":{\"btntxt\":\"\",\"description\":\"描述\",\"title\":\"标题\",\"url\":\"url\"},\"touser\":\"test\"}")
}

func TestMessage_Markdown(t *testing.T) {
	m := msg.Clone().ToUser([]string{"test"}).Markdown("您的会议室已经预定")

	assertEqual(t,
		m.ToJson(),
		"{\"agentid\":123,\"markdown\":{\"content\":\"您的会议室已经预定\"},\"msgtype\":\"markdown\",\"touser\":\"test\"}")
}
