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
		AccountSid: "xxxxxxxxxxx",
		AuthToken:  "xxxxxxxxxxx",
		MsgData:    url.Values{},
	}
	t.UrlStr = "https://api.twilio.com/2010-04-01/Accounts/" + t.AccountSid + "/Messages.json"
	t.MsgData.Set("To", "pppppppppp")
	t.MsgData.Set("From", "pppppppppp")
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
