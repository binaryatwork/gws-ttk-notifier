package twilio

import (
	"errors"
	"fmt"
	"github.com/carnellj/gws-ttk-notifier/utils"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type TwilioTextNotify struct {
	Config *utils.Config `inject:""`
}

func (t *TwilioTextNotify) Notify(Target string, Message string) (StatusCode int, Error error) {
	log.Printf("TWILIO_AUTH: %s", t.Config.TwilioAuth)
	log.Printf( "TWILIO_SID: %s", t.Config.TwilioSid)
	log.Printf("TWILIO_NUMBER: %s", t.Config.TwilioNumber)
	twilioUrl := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages", t.Config.TwilioSid)

	form := url.Values{}
	form.Add("Body", Message)
	form.Add("From", t.Config.TwilioNumber)
	form.Add("To", Target)

	req, _ := http.NewRequest("POST", twilioUrl, strings.NewReader(form.Encode()))
	req.SetBasicAuth(t.Config.TwilioSid, t.Config.TwilioAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	log.Printf("TWILIO RESPONSE: %s\n", err)
	if err != nil {
		log.Printf("Error encountered while calling twilio api: %s\n", err.Error())
		return http.StatusBadRequest, err
	}

	if resp.StatusCode > 299 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		log.Printf("The call to twilio return a error statusCode %d with error %s\n", resp.StatusCode, bodyString)
		return resp.StatusCode, errors.New(string(bodyString))
	}

	return resp.StatusCode, nil
}
