package twilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type TwilioAPI struct {
	AccountSid string
	AuthToken  string
	UrlStr     string
	MsgData    url.Values
}

func NewTwilioAPI() *TwilioAPI {
	t := &TwilioAPI{
		// Set account keys & information
		AccountSid: "AC0f5b6938d0dd092d0823f8fbccda30f7",
		AuthToken:  "3ef748eafff2c62e78843f462231e2e5",
		MsgData:    url.Values{},
	}
	t.UrlStr = "https://api.twilio.com/2010-04-01/Accounts/" + t.AccountSid + "/Messages.json"
	t.MsgData.Set("To", "+2349086790286")
	t.MsgData.Set("From", "(909) 639-4683")
	return t
}

func (t *TwilioAPI) TwilioTextMessage(TextMessage string) {
	// Build out the data for our message
	t.MsgData.Set("Body", TextMessage)
	msgDataReader := *strings.NewReader(t.MsgData.Encode())
	// Create Client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", t.UrlStr, &msgDataReader)
	req.SetBasicAuth(t.AccountSid, t.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	twilioResp, _ := client.Do(req)
	if twilioResp.StatusCode >= 200 && twilioResp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(twilioResp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(twilioResp.Status)
	}
}
