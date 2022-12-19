package utils

import "fmt"

func IsValidContent(content string) error {
	// Check the content size less than 5
	if len(content) < 5 {
		return fmt.Errorf("The message content must be at least 5 characters long")
	}
	return nil
}

