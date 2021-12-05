package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	vp *viper.Viper
}

func init() {
	err := setupConfig()
	if err != nil {
		log.Fatalf("init.setupConfig err: %v", err)
	}
}

func setupConfig() error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}
	err = config.ReadSection("Mysql", &MysqlC)
	if err != nil {
		return err
	}
	err = config.ReadSection("Server", &ServerC)
	if err != nil {
		return err
	}
	err = config.ReadSection("nsq", &NsqC)

	return nil
}

func LoadConfig() (*Config, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("config/")
	vp.SetConfigType("json")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{vp: vp}, nil
}
