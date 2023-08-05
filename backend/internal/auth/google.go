package auth

import (
	"context"
	"sync"

	"github.com/ak9024/okr-generator/config"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauth2_v2 "google.golang.org/api/oauth2/v2"
)

var (
	once                sync.Once
	authConfig          *oauth2.Config
	state               string
	googleAuthLogoutURL = "https://accounts.google.com/logout"
	googleScopes        = []string{"email", "profile"}
)

// Setup auth with Google login
func getConfig(cfg config.Provider) *oauth2.Config {
	once.Do(func() {
		authConfig = &oauth2.Config{
			RedirectURL:  cfg.GetString("google.redirect_url"),
			ClientID:     cfg.GetString("google.client_id"),
			ClientSecret: cfg.GetString("google.client_secret"),
			Scopes:       googleScopes,
			Endpoint:     google.Endpoint,
		}
	})

	return authConfig
}

func generateRedirectUrl(cfg config.Provider) string {
	config := getConfig(cfg)
	// generate random key as a state
	state = uuid.New().String()
	// generate redirect url
	redirectUrl := config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return redirectUrl
}

func generateGoogleToken(c *fiber.Ctx, cfg config.Provider) (*oauth2.Token, error) {
	config := getConfig(cfg)

	// oauth exchange
	token, errToken := config.Exchange(context.Background(), c.Query("code"))
	if errToken != nil {
		return nil, errToken
	}

	return token, nil
}

func initGoogleClient(cfg config.Provider, token *oauth2.Token) (*oauth2_v2.Service, error) {
	config := getConfig(cfg)

	client := config.Client(context.Background(), token)

	service, errService := oauth2_v2.New(client)
	if errService != nil {
		return nil, errService
	}

	return service, nil
}
