package messenger

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCreativeMessageCreative(t *testing.T) {
	//Avoid HTTPS in tests
	GraphAPI = "http://example.com"
	messenger := &Messenger{}

	mockData := &MessageCreativeResponse{
		MessageCreativeID: 1,
	}
	body, err := json.Marshal(mockData)
	if err != nil {
		t.Error(err)
	}

	setClient(200, body)

	mcr := MessageCreativeRequest{[]SendMessage{SendMessage{
		Text: "Hello World",
	}}}
	response, err := messenger.CreateMessageCreative(mcr)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(response, mockData) {
		t.Error("response do not match")
	}

	errorData := &rawError{Error: Error{
		Message: "w/e",
	}}
	body, err = json.Marshal(errorData)
	if err != nil {
		t.Error(err)
	}
	setClient(400, body)
	mcr = MessageCreativeRequest{}
	_, err = messenger.CreateMessageCreative(mcr)
	if err == nil {
		t.Error("Error shouldn't be empty.")
	}
}

func TestSendBroadcast(t *testing.T) {
	//Avoid HTTPS in tests
	GraphAPI = "http://example.com"
	messenger := &Messenger{}

	mockData := &BroadcastResponse{
		BroadcastID: 1,
	}
	body, err := json.Marshal(mockData)
	if err != nil {
		t.Error(err)
	}

	setClient(200, body)

	request := BroadcastRequest{
		MessageCreativeId: 1,
		MessagingType:     MessagingTypeTag,
		MessageTag:        MessageTagNonPromotionalSubscription,
	}
	response, err := messenger.SendBroadcast(request)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(response, mockData) {
		t.Error("response do not match")
	}

	errorData := &rawError{Error: Error{
		Message: "w/e",
	}}
	body, err = json.Marshal(errorData)
	if err != nil {
		t.Error(err)
	}
	setClient(400, body)
	_, err = messenger.SendBroadcast(request)
	if err == nil {
		t.Error("Error shouldn't be empty.")
	}

	request = BroadcastRequest{
		MessageCreativeId: 1,
		MessageTag:        MessageTagNonPromotionalSubscription,
	}
	_, err = messenger.SendBroadcast(request)
	if err.Error() != "Messaging Type must be set to MESSAGE_TAG" {
		t.Error("Messaging type should be set")
	}

	request = BroadcastRequest{
		MessageCreativeId: 1,
		MessagingType:     MessagingTypeTag,
	}
	_, err = messenger.SendBroadcast(request)
	if err.Error() != "Notification type Must be set to NON_PROMOTIONAL_SUBSCRIPTION" {
		t.Error("Messaging tag should be set")
	}
}
