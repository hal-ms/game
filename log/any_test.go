package log

import "testing"

func TestSendSlack(t *testing.T) {
	SendSlack("hello!")
}
