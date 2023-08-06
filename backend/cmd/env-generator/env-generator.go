package envgenerator

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/sirupsen/logrus"
)

type EnvGenerator struct {
	Port                    int
	Host                    string
	Version                 string
	Env                     string
	Token                   string
	GoogleClientID          string
	GoogleClientSecret      string
	GoogleRedirectURL       string
	GoogleClientRedirectURL string
	SupabaseURL             string
	SupabaseKey             string
}

func New(cfg EnvGenerator) *EnvGenerator {
	return &EnvGenerator{
		Port:                    cfg.Port,
		Host:                    cfg.Host,
		Version:                 cfg.Version,
		Env:                     cfg.Env,
		Token:                   cfg.Token,
		GoogleClientID:          cfg.GoogleClientID,
		GoogleClientSecret:      cfg.GoogleClientSecret,
		GoogleRedirectURL:       cfg.GoogleRedirectURL,
		GoogleClientRedirectURL: cfg.GoogleClientRedirectURL,
		SupabaseURL:             cfg.SupabaseURL,
		SupabaseKey:             cfg.SupabaseKey,
	}
}

func (eg *EnvGenerator) ConvertEnvIntoYaml() []byte {
	cfg := map[string]interface{}{
		"app": map[string]interface{}{
			"port":    eg.Port,
			"host":    eg.Host,
			"version": eg.Version,
			"env":     eg.Env,
		},
		"chatgpt": map[string]interface{}{
			"token": eg.Token,
		},
		"google": map[string]interface{}{
			"redirect_url":        eg.GoogleRedirectURL,
			"client_id":           eg.GoogleClientID,
			"client_secret":       eg.GoogleClientSecret,
			"client_redirect_url": eg.GoogleClientRedirectURL,
		},
		"supabase": map[string]interface{}{
			"url": eg.SupabaseURL,
			"key": eg.SupabaseKey,
		},
	}

	// Convert the data to toml structure
	b, errCreate := toml.Marshal(cfg)
	if errCreate != nil {
		logrus.Error(errCreate)
	}

	return b
}

func (eg *EnvGenerator) Exec() {
	// Convert env into file with toml format
	b := eg.ConvertEnvIntoYaml()

	// Create a file .config.generated.toml
	f, errGenerateFile := os.Create(".config.generated.toml")
	if errGenerateFile != nil {
		logrus.Error(errGenerateFile)
	}
	defer f.Close()

	// Insert the data to file `.config.generated.toml`
	_, errWrite := f.Write(b)
	if errWrite != nil {
		logrus.Error(errWrite)
	}

	logrus.Info(fmt.Sprintf("Success to generate: %v", f.Name()))
}
