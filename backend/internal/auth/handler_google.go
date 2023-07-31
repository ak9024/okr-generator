package auth

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	oauth2_v2 "google.golang.org/api/oauth2/v2"
)

func (a *auth) GoogleLoginHandler(c *fiber.Ctx) error {
	url := authConfig.AuthCodeURL(state)
	return c.Redirect(url, http.StatusTemporaryRedirect)
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

	if userInfo != nil {
		emailInDB, _ := a.ViewUserFilterByEmail(userInfo.Email)
		if emailInDB == nil {
			um := UserModel{
				UUID:    uuid.New(),
				Name:    userInfo.Name,
				Email:   userInfo.Email,
				EmailID: userInfo.Id,
				Picture: userInfo.Picture,
			}

			_, errInsert := a.InsertUser(um)
			if errInsert != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		response := GoogleLoginCallbackResponse200{
			StatusCode: fiber.StatusOK,
			User: User{
				ID:            userInfo.Id,
				Name:          userInfo.Name,
				Email:         userInfo.Email,
				VerifiedEmail: *userInfo.VerifiedEmail,
				Token:         token.AccessToken,
				FamilyName:    userInfo.FamilyName,
				GivenName:     userInfo.GivenName,
				Locale:        userInfo.Locale,
				Picture:       userInfo.Picture,
			},
		}

		return c.Status(fiber.StatusOK).JSON(response)
	}

	return c.SendStatus(fiber.StatusBadRequest)
}

func (a *auth) GoogleLogoutHandler(c *fiber.Ctx) error {
	return c.Redirect(googleAuthLogoutURL, http.StatusTemporaryRedirect)
}
