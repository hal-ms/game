package log

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

var SlackURL = "https://hooks.slack.com/services/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"

func SendSlack(msg string) {
	payload := map[string]interface{}{
		"text": msg,
	}
	str, err := json.Marshal(payload)
	if err != nil {
		return
	}
	values := url.Values{}
	values.Set("payload", string(str))

	req, _ := http.NewRequest(
		"POST",
		SlackURL,
		strings.NewReader(values.Encode()),
	)

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	client.Do(req)
}
