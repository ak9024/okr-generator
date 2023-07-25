package service

import "github.com/ak9024/okr-generator/config"

type Service struct {
	config config.Provider
}

func NewService(cfg config.Provider) *Service {
	return &Service{
		config: cfg,
	}
}
