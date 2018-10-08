package messenger

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// The id of  the Inbox, useful when you want to pass your control to a human !
const InboxID = 263902037430900

type passThreadControlRequest struct {
	RecipientRequest recipientRequest `json:"recipient,omitempty"`
	TargetAppID      int              `json:"target_app_id,omitempty"`
	Metadata         string           `json:"metadata,omitempty"`
}

type threadControlRequest struct {
	RecipientRequest recipientRequest `json:"recipient_request,omitempty"`
	Metadata         string           `json:"metadata,omitempty"`
}

type recipientRequest struct {
	ID int `json:"id,omitempty"`
}

func (m *Messenger) CallPassThreadControl(recipientID int, appID int, metadata string) ([]byte, error) {

	request := passThreadControlRequest{
		RecipientRequest: recipientRequest{
			ID: recipientID,
		},
		TargetAppID: appID,
		Metadata:    metadata,
	}

	byt, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := m.doRequest(http.MethodPost, GraphAPI+"/v3.1/me/pass_thread_control", bytes.NewReader(byt))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		er := new(rawError)
		json.Unmarshal(read, er)
		return nil, er.Error
	}
	return read, err
}

func (m *Messenger) CallTakeThreadControl(recipientID int, metadata string) ([]byte, error) {
	return m.requestThreadControl(recipientID, metadata, "/v3.1/me/take_thread_control")
}

func (m *Messenger) CallRequestThreadControl(recipientID int, metadata string) ([]byte, error) {
	return m.requestThreadControl(recipientID, metadata, "/v3.1/me/request_thread_control")
}

func (m *Messenger) requestThreadControl(recipientID int, metadata string, uri string) ([]byte, error) {
	request := threadControlRequest{
		RecipientRequest: recipientRequest{
			ID: recipientID,
		},
		Metadata: metadata,
	}
	byt, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := m.doRequest(http.MethodPost, GraphAPI+uri, bytes.NewReader(byt))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		er := new(rawError)
		json.Unmarshal(read, er)
		return nil, er.Error
	}
	return read, err
}

func (m *Messenger) GetSecondaryReceivers() ([]byte, error) {
	resp, err := m.doRequest(http.MethodPost, GraphAPI+"/v3.1/me/secondary_receivers?fields=id,name", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		er := new(rawError)
		json.Unmarshal(read, er)
		return nil, er.Error
	}
	return read, err
}
