package auth

import "github.com/ak9024/okr-generator/config"

type auth struct {
	Config config.Provider
}

func NewAuth(cfg config.Provider) *auth {
	return &auth{
		Config: cfg,
	}
}
