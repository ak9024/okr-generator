package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/ak9024/okr-generator/cmd/server"
	"github.com/ak9024/okr-generator/config"
	"github.com/ak9024/okr-generator/internal/okr"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGoogleLogin(t *testing.T) {
	t.Log("GET /api/auth/google/login expected 307 redirect")
	cfg := config.Config()
	app := server.NewServer(cfg)
	router := app.Router()

	req := httptest.NewRequest(fiber.MethodGet, "/api/auth/google/login", nil)

	resp, err := router.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusTemporaryRedirect, resp.StatusCode)
}

func TestGoogleCallback(t *testing.T) {
	t.Log("GET /api/auth/google/callback expected 400 redirect")
	cfg := config.Config()
	app := server.NewServer(cfg)
	router := app.Router()

	req := httptest.NewRequest(fiber.MethodGet, "/api/auth/google/callback", nil)

	resp, err := router.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestGoogleLogout(t *testing.T) {
	t.Log("GET /api/auth/google/logout expected 307 redirect")
	cfg := config.Config()
	app := server.NewServer(cfg)
	router := app.Router()

	req := httptest.NewRequest(fiber.MethodGet, "/api/auth/google/logout", nil)

	resp, err := router.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusTemporaryRedirect, resp.StatusCode)
}

func TestSubmitOKR(t *testing.T) {
	t.Log("POST /api/v1/okr-generator expected 401 redirect")
	cfg := config.Config()
	app := server.NewServer(cfg)
	router := app.Router()

	payload := okr.OKRGeneratorRequest{
		Objective: "Test objective submit",
		Translate: "Bahasa",
	}

	jsonMarshal, err := json.Marshal(payload)
	assert.Nil(t, err)

	req := httptest.NewRequest(fiber.MethodPost, "/api/v1/okr-generator", bytes.NewBuffer(jsonMarshal))
	req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	req.Header.Add(fiber.HeaderAuthorization, "Bearer token-here")

	resp, err := router.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
}
