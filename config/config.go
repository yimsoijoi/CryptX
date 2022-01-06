package config

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/yimsoijoi/cryptx/lib/postgres"
)

type Config struct {
	WalletPrivateKey string          `mapstructure:"wallet"`
	Postgres         postgres.Config `mapstructure:"postgres"`
}

func Load() (*Config, error) {
	var conf Config
	// /dir/name.ext

	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := viper.SafeWriteConfig(); err != nil {
				log.Println("can't write new config")
				return nil, errors.Wrap(err, "can't write config")
			}
		}
		return nil, errors.Wrap(err, "can't read config")
	}

	if err := viper.Unmarshal(&conf); err != nil {
		log.Println("can't unmarshal config")
	}

	return &conf, nil
}
