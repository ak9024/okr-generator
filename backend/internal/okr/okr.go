package okr

import "github.com/ak9024/okr-generator/config"

type okr struct {
	Config config.Provider
}

func NewOKR(cfg config.Provider) *okr {
	return &okr{
		Config: cfg,
	}
}
