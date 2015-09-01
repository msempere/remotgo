package utils

import (
	"fmt"
	"os/user"
)

// GetUsername return the current username for the running user
// or empty string in case error
func GetUsername() string {
	username, err := user.Current()
	if err != nil {
		return ""
	}
	return username.Username
}

func GetDefaultRSAFilePath() string {
	username := GetUsername()

	if len(username) == 0 {
		return ""
	}

	userentry, err := user.Lookup(username)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s/.ssh/id_rsa", userentry.HomeDir)
}

func GetDefaultDSAFilePath() string {
	username := GetUsername()

	if len(username) == 0 {
		return ""
	}
	userentry, err := user.Lookup(username)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s/.ssh/id_dsa", userentry.HomeDir)
}
