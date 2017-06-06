package freemail

import "testing"

func TestGmail(t *testing.T) {
	gmail := "user@gmail.com"
	if !IsFreemail(gmail) {
		t.Fail()
	}
}
