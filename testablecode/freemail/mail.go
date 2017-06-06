package freemail

import "strings"

var freemails = []string{"gmail.com", "yahoo.com", "outlook.com"}

// IsFreemail tests the email address for known TLD suffixes and returns true if there is a match
func IsFreemail(email string) bool {
	for _, provider := range freemails {
		if strings.Contains(email, provider) {
			return true
		}
	}
	return false
}
