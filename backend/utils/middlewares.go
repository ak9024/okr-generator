package utils

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware(c *fiber.Ctx) to prevent user access if not have credential
func AuthMiddleware(c *fiber.Ctx) error {
	// Get token from header "Authorization"
	authorization := c.Get("Authorization")
	token := strings.SplitN(authorization, " ", -1)[1]

	// Get data from user profile
	resp, _ := GetUserProfile(token)
	// if resp.ID is not empty approved the request
	if resp.ID != "" {
		return c.Next()
	}

	// if not please sent unauthorized satus code
	return c.SendStatus(fiber.StatusUnauthorized)
}

// GetUserProfile(token string) to validate the token in google profile
func GetUserProfile(token string) (*UserProfileResponse200, *UserProfileReponseError) {
	url := fmt.Sprintf("%s?access_token=%s", endpointProfile, url.QueryEscape(token))

	result := UserProfileResponse200{}
	errResult := UserProfileReponseError{}

	client := resty.New()
	_, err := client.R().
		SetResult(&result).
		SetError(&errResult).
		Get(url)
	if err != nil {
		return nil, &errResult
	}

	return &result, nil
}
