package auth

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	oauth2_v2 "google.golang.org/api/oauth2/v2"
)

func (a *auth) GoogleLoginHandler(c *fiber.Ctx) error {
	url := authConfig.AuthCodeURL(state)
	return c.Redirect(url, http.StatusTemporaryRedirect)
}

func (a *auth) GoogleLogoutHandler(c *fiber.Ctx) error {
	return c.Redirect(googleAuthLogoutURL, http.StatusTemporaryRedirect)
}

func (a *auth) GoogleLoginCallback(c *fiber.Ctx) error {
	token, errToken := authConfig.Exchange(context.Background(), c.FormValue("code"))
	if errToken != nil {
		c.SendStatus(fiber.StatusBadRequest)
	}

	client := authConfig.Client(context.Background(), token)

	service, errService := oauth2_v2.New(client)
	if errService != nil {
		c.SendStatus(fiber.StatusBadRequest)
	}

	userInfo, errGetUserInfo := service.Userinfo.Get().Do()
	if errGetUserInfo != nil {
		c.SendStatus(fiber.StatusBadRequest)
	}

	response := GoogleLoginCallbackResponse200{
		StatusCode: fiber.StatusOK,
		UserInfo:   userInfo,
		Token:      token,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
