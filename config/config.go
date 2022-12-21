package config

import (
	"fmt"
	"path/filepath"
)
import "github.com/spf13/viper"

type AppConfig struct {
	Debug           bool     `json:"debug"`
	Port            int      `json:"port"`
	AutoCreateUsers []string `json:"auto_create_users" mapstructure:"auto_create_users"`
}

func (receiver *AppConfig) ReInit(name string) error  {
	setup(name)
	err := viper.Unmarshal(receiver)
	return err
}
func NewConfig() *AppConfig {
	setup("config.json")
	var conf AppConfig
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	return &conf
}



func setup(config string) {
	extension := filepath.Ext(config)
	var name = config[0 : len(config)-len(extension)]
	viper.SetConfigName(name)           // name of config file (without extension)
	viper.SetConfigType(extension[1:])  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")            // optionally look for config in the working directory
	viper.AddConfigPath("/etc/oplin/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.oplin") // call multiple times to add many search paths
	err := viper.ReadInConfig()         // Find and read the config file
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

}
