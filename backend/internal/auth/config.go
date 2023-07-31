package auth

import (
	"github.com/ak9024/okr-generator/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Setup auth with Google login
func SetupOauth(cfg config.Provider) {
	authConfig = &oauth2.Config{
		RedirectURL:  cfg.GetString("google.redirect_url"),
		ClientID:     cfg.GetString("google.client_id"),
		ClientSecret: cfg.GetString("google.client_secret"),
		Scopes:       googleScopes,
		Endpoint:     google.Endpoint,
	}

	state = cfg.GetString("google.auth_state")
}
