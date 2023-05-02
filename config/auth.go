package config

import (
	"time"

	"github.com/paycrest/paycrest-protocol/utils/logger"
	"github.com/spf13/viper"
)

type AuthConfiguration struct {
	Secret                 string
	JwtAccessHourLifespan  time.Duration
	JwtRefreshHourLifespan time.Duration
}

func AuthConfig() (config *AuthConfiguration) {

	return &AuthConfiguration{
		Secret:                 viper.GetString("SECRET"),
		JwtAccessHourLifespan:  time.Duration(15) * time.Minute,
		JwtRefreshHourLifespan: time.Duration(24) * time.Hour,
	}
}

func init() {
	if err := SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}
}
