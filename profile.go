package messenger

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Profile struct holds data associated with Facebook profile
type Profile struct {
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	ProfilePicture string  `json:"profile_pic,omitempty"`
	Locale         string  `json:"locale,omitempty"`
	Timezone       float64 `json:"timezone,omitempty"`
	Gender         string  `json:"gender,omitempty"`
}

// GetProfile fetches the recipient's profile from facebook platform
// Non empty UserID has to be specified in order to receive the information
func (m *Messenger) GetProfile(userID string) (*Profile, error) {
	return m.fetchProfile(fmt.Sprintf(GraphAPI+"/v3.1/%s?fields=first_name,last_name,profile_pic", userID))
}

func (m *Messenger) GetProfileWithLocale(userID string) (*Profile, error) {
	return m.fetchProfile(fmt.Sprintf(GraphAPI+"/v3.1/%s?fields=first_name,last_name,profile_pic,locale", userID))
}

func (m *Messenger) GetProfileWithTimeZone(userID string) (*Profile, error) {
	return m.fetchProfile(fmt.Sprintf(GraphAPI+"/v3.1/%s?fields=first_name,last_name,profile_pic,timezone", userID))
}

func (m *Messenger) GetProfileWithGender(userID string) (*Profile, error) {
	return m.fetchProfile(fmt.Sprintf(GraphAPI+"/v3.1/%s?fields=first_name,last_name,profile_pic,gender", userID))
}

func (m *Messenger) fetchProfile(profileEndpoint string) (*Profile, error) {
	resp, err := m.doRequest("GET", profileEndpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		er := new(rawError)
		json.Unmarshal(read, er)
		return nil, errors.New("Error occured: " + er.Error.Message)
	}
	profile := new(Profile)
	return profile, json.Unmarshal(read, profile)
}

type accountLinking struct {
	//Recipient is Page Scoped ID
	Recipient string `json:"recipient"`
}

// GetPSID fetches user's page scoped id during authentication flow
// one must supply a valid and not expired authentication token provided by facebook
// https://developers.facebook.com/docs/messenger-platform/account-linking/authentication
func (m *Messenger) GetPSID(token string) (*string, error) {
	resp, err := m.doRequest("GET", fmt.Sprintf(GraphAPI+"/v2.6/me?fields=recipient&account_linking_token=%s", token), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		er := new(rawError)
		json.Unmarshal(read, er)
		return nil, errors.New("Error occured: " + er.Error.Message)
	}
	acc := new(accountLinking)
	return &acc.Recipient, json.Unmarshal(read, acc)
}
