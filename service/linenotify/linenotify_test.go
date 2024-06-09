package linenotify

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLINENotifySend(t *testing.T) {
	n := Notification{
		Token:                "token",
		Message:              "Testing notification",
		NotificationDisabled: false,
	}

	var mockResp apiResponse
	var hitServer bool

	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		hitServer = true

		if r.Method != "POST" {
			t.Error("HTTP method should be POST.")
		}

		if n.Token == "" {
			t.Error("missing token")
		}

		if n.Message == "" {
			t.Error("missing message")
		}

		if err := json.NewEncoder(rw).Encode(mockResp); err != nil {
			t.Error(err)
		}
	}))
	defer ts.Close()

	API = ts.URL

	// successful
	mockResp.Status = 200
	if err := n.Send(); err != nil {
		t.Error(err)
	}

	if !hitServer {
		t.Error("didn't reach server")
	}
}
