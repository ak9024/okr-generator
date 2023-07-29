package auth

import (
	"github.com/ak9024/okr-generator/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfig *oauth2.Config
	authState   string
)

func SetupOauth(cfg config.Provider) {
	oauthConfig = &oauth2.Config{
		RedirectURL:  cfg.GetString("google.redirect_url"),
		ClientID:     cfg.GetString("google.client_id"),
		ClientSecret: cfg.GetString("google.client_secret"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	authState = cfg.GetString("google.auth_state")
}
