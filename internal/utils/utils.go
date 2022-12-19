package utils

import "fmt"

var (
	ignoreUser = []string{
		// "623155071735037982",
		"983924510220779550",
	}
	ignoreChannel = []string{
		"1021958640829210674",
	}
)

func IsValidContent(content string) error {
	// Check the content size less than 5
	if len(content) < 5 {
		return fmt.Errorf("The message content must be at least 5 characters long")
	}
	return nil
}

func IgnoreUser(userID string) bool {
	for _, u := range ignoreUser {
		if u == userID {
			return true
		}
	}
	return false
}
