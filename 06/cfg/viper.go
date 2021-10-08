package cfg

import (
	"github.com/spf13/viper"
	"log"
)

func SetupConfig() {
	viper.AddConfigPath(".")

	viper.SetConfigName("config.default") // config.toml
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Config file not found")
		} else {
			log.Fatalf("Errorf reading conf file %s", err.Error())
		}
	}
}
