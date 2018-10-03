package messenger

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type MessageCreativeRequest struct {
	Messages []SendMessage `json:"messages,omitempty"`
}

type MessageCreativeResponse struct {
	MessageCreativeID int `json:"message_creative_id,string,omitempty"`
}

type BroadcastRequest struct {
	MessageCreativeId int              `json:"message_creative_id,omitempty"`
	NotificationType  NotificationType `json:"notification_type,omitempty"`
	MessagingType     MessagingType    `json:"messaging_type,omitempty"`
	MessageTag        MessageTag       `json:"message_tag,omitempty"`
}

type BroadcastResponse struct {
	BroadcastID int `json:"broadcast_id,string,omitempty"`
}

func (m *Messenger) CreateMessageCreative(mcr MessageCreativeRequest) (*MessageCreativeResponse, error) {
	byt, err := json.Marshal(mcr)
	if err != nil {
		return nil, err
	}
	resp, err := m.doRequest("POST", GraphAPI+"/v3.1/me/message_creatives", bytes.NewReader(byt))
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
	response := new(MessageCreativeResponse)
	return response, json.Unmarshal(read, response)
}

func (m *Messenger) SendBroadcast(br BroadcastRequest) (*BroadcastResponse, error) {

	if br.MessagingType != MessagingTypeTag {
		return nil, errors.New("Messaging Type must be set to MESSAGE_TAG")
	}
	if br.MessageTag != MessageTagNonPromotionalSubscription {
		return nil, errors.New("Notification type Must be set to NON_PROMOTIONAL_SUBSCRIPTION")
	}
	byt, err := json.Marshal(br)
	if err != nil {
		return nil, err
	}
	resp, err := m.doRequest("POST", GraphAPI+"/v3.1/me/broadcast_messages", bytes.NewReader(byt))
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
	response := new(BroadcastResponse)
	return response, json.Unmarshal(read, response)
}
