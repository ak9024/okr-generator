package auth

import (
	"golang.org/x/oauth2"
)

var (
	authConfig *oauth2.Config
	state      string
)
