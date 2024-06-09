package linenotify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var API = "https://notify-api.line.me/api/notify"

type apiResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Notification struct {
	// Token is the user's authentication token.
	Token string
	// Message is the notification's message.
	Message string
	// Wheteher to disable notification.
	NotificationDisabled bool
}

func (n *Notification) Send() error {
	if n.Token == "" {
		return errors.New("missing token")
	}

	if n.Message == "" {
		return errors.New("missing message")
	}

	vals := make(url.Values)
	vals.Set("message", n.Message)
	vals.Set("notificationDisabled", fmt.Sprintf("%t", n.NotificationDisabled))
	valsReader := bytes.NewBufferString(vals.Encode())

	client := &http.Client{}
	req, err := http.NewRequest("POST", API, valsReader)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", n.Token))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var r apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return err
	}

	if r.Status != 200 {
		return errors.New(r.Message)
	}

	return nil
}
