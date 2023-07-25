package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/integralist/go-findroot/find"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	defaultConfig *viper.Viper
	once          sync.Once
)

type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

func readViperConfig() *viper.Viper {
	basePath := GetBasePath()

	v := viper.New()
	v.SetConfigName(".config")
	v.AddConfigPath(".")
	v.AddConfigPath(fmt.Sprintf("%s/", basePath))

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Errorf("Error: %s", err)
		} else {
			panic(fmt.Errorf("Config error: %s", err))
		}
	}

	fmt.Printf("Using config file: %s \n\n", v.ConfigFileUsed())

	return v
}

func Config() Provider {
	once.Do(func() {
		defaultConfig = readViperConfig()
	})

	return defaultConfig
}

func GetBasePath() string {
	root, err := find.Repo()
	if err != nil {
		logrus.Errorf("Error: %s", err.Error())
	}

	return root.Path
}
