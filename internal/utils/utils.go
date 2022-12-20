package utils

import "fmt"

var (
	ignoreUser = []string{
		// "623155071735037982", me
		"983924510220779550",  // wen
		"1026733912778625026", // corrie
		"912897330213179402",  // rosie
		"885891259053531176",  // semi
		"948825318515425280",  // sky
	}
	ignoreChannel = []string{
		"1021958640829210674", // test server (attendance)
		"808621206718251058",  // moderator-only
		"537522976963166218",  // announcements.
		"583944383083184129",  // playdapp-sns.
		"570896878858665984",  // welcome.
		"583944743655047178",  // rules-and-admin-team.
		"920238004147204177",  // filipino.
		"585672690111610880",  // chinese.
		"585672615683686419",  // russian.
		"583934248512258059",  // japanese.
		"585672591449260032",  // vietnamese.
		"1016194558926803075", // indonesia
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

func IgnoreChannel(channelID string) bool {
	for _, c := range ignoreChannel {
		if c == channelID {
			return true
		}
	}
	return false
}
