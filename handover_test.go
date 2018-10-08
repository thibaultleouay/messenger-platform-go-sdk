package messenger

import "testing"

func TestCallPassThread(t *testing.T) {
	GraphAPI = "http://example.com"
	messenger := &Messenger{}

	str := `{"success":true}`

	setClient(200, []byte(str))

	messenger.CallPassThreadControl(1, InboxID, "")
}

func TestCallTakeThread(t *testing.T) {
	GraphAPI = "http://example.com"
	messenger := &Messenger{}

	str := `{"success":true}`

	setClient(200, []byte(str))

	messenger.CallTakeThreadControl(1, "")
}

func TestCallRequestThread(t *testing.T) {
	GraphAPI = "http://example.com"
	messenger := &Messenger{}

	str := `{"success":true}`

	setClient(200, []byte(str))

	messenger.CallRequestThreadControl(1, "")
}
