package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (a *auth) GoogleLoginHandler(c *fiber.Ctx) error {
	redirectUrl := generateRedirectUrl(a.Config)
	return c.Redirect(redirectUrl, fiber.StatusTemporaryRedirect)
}

func (a *auth) GoogleLoginCallback(c *fiber.Ctx) error {
	token, errToken := generateGoogleToken(c, a.Config)
	if errToken != nil {
		c.SendStatus(fiber.StatusBadRequest)
	}

	service, errService := initGoogleClient(a.Config, token)
	if errService != nil {
		c.SendStatus(fiber.StatusInternalServerError)
	}

	// get data profile
	userInfo, errGetUserInfo := service.Userinfo.Get().Do()
	if errGetUserInfo != nil {
		c.SendStatus(fiber.StatusInternalServerError)
	}

	// userInfo ready, next process the data
	if userInfo != nil {
		// check the email in supabase storage
		emailInDB, _ := a.ViewUserFilterByEmail(userInfo.Email)
		// if email not exists, insert the data as a new email
		if emailInDB == nil {
			um := UserModel{
				UUID:    uuid.New(),
				Name:    userInfo.Name,
				Email:   userInfo.Email,
				EmailID: userInfo.Id,
				Picture: userInfo.Picture,
			}

			// insert new email
			_, errInsert := a.InsertUser(um)
			if errInsert != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		c.Cookie(&fiber.Cookie{
			Name:    "token",
			Value:   token.AccessToken,
			Expires: token.Expiry,
			Domain:  a.Config.GetString("google.client_redirect_url"),
		})

		return c.Redirect(a.Config.GetString("google.client_redirect_url"), fiber.StatusTemporaryRedirect)
	}

	return c.SendStatus(fiber.StatusBadRequest)
}

func (a *auth) GoogleLogoutHandler(c *fiber.Ctx) error {
	return c.Redirect(googleAuthLogoutURL, http.StatusTemporaryRedirect)
}
