package auth

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	oauth2_v2 "google.golang.org/api/oauth2/v2"
)

func (a *auth) GoogleLoginHandler(c *fiber.Ctx) error {
	url := oauthConfig.AuthCodeURL(authState)
	return c.Redirect(url, http.StatusTemporaryRedirect)
}

func (a *auth) GoogleLoginCallback(c *fiber.Ctx) error {
	token, _ := oauthConfig.Exchange(context.Background(), c.FormValue("code"))
	client := oauthConfig.Client(context.Background(), token)
	service, _ := oauth2_v2.New(client)
	userInfo, _ := service.Userinfo.Get().Do()

	response := LoginResponse200{
		StatusCode: fiber.StatusOK,
		Data:       userInfo,
		Token:      token,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
