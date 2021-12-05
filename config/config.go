package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	vp *viper.Viper
}

func setupConfig(config *Config) error {
	err := config.ReadSection("Mysql", &MysqlC)
	if err != nil {
		return err
	}
	err = config.ReadSection("Server", &ServerC)
	if err != nil {
		return err
	}
	err = config.ReadSection("nsq", &NsqC)
	if err != nil {
		return err
	}

	return nil
}

func LoadConfig(root string) error {
	vp := viper.New()
	vp.SetEnvPrefix(root)
	vp.SetConfigName("config")
	vp.AddConfigPath(fmt.Sprintf("%sconfig/", root))
	vp.SetConfigType("json")
	err := vp.ReadInConfig()
	if err != nil {
		return err
	}

	config := &Config{vp: vp}
	err = setupConfig(config)
	if err != nil {
		return err
	}
	return nil
}
