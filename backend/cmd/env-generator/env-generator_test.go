package envgenerator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvGenerator(t *testing.T) {
	t.Log("Test env generator config")

	eg := New(EnvGenerator{
		Port:                    3000,
		Host:                    "http://localhost",
		Version:                 "v1.x.x",
		Env:                     "testing",
		Token:                   "token",
		GoogleClientID:          "google-client-id",
		GoogleClientSecret:      "google-client-secret",
		GoogleRedirectURL:       "google-redirect-url",
		GoogleClientRedirectURL: "google-client-redirect-url",
		SupabaseURL:             "supabase-url",
		SupabaseKey:             "supabase-key",
	})

	assert.Equal(t, 3000, eg.Port)
	assert.Equal(t, "http://localhost", eg.Host)
	assert.Equal(t, "v1.x.x", eg.Version)
	assert.Equal(t, "testing", eg.Env)
	assert.Equal(t, "token", eg.Token)
	assert.Equal(t, "google-client-id", eg.GoogleClientID)
	assert.Equal(t, "google-client-secret", eg.GoogleClientSecret)
	assert.Equal(t, "google-redirect-url", eg.GoogleRedirectURL)
	assert.Equal(t, "google-client-redirect-url", eg.GoogleClientRedirectURL)
	assert.Equal(t, "supabase-url", eg.SupabaseURL)
	assert.Equal(t, "supabase-key", eg.SupabaseKey)

	b := eg.ConvertEnvIntoYaml()
	t.Log(string(b))
}
