package auth

import (
	"golang.org/x/oauth2"
)

var (
	authConfig          *oauth2.Config
	state               string
	googleAuthLogoutURL = "https://accounts.google.com/logout"
	googleScopes        = []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	}
)
