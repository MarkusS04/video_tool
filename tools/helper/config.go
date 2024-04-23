package helper

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type config struct {
	JW struct {
		ApiUrl string
		All    string
	}
	FS struct {
		Source      string
		Destination string
	}
}

var Config config

func GetConfig() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	var conf config
	if err := viper.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	conf.FS.Source = os.ExpandEnv(conf.FS.Source)
	conf.FS.Destination = os.ExpandEnv(conf.FS.Destination)

	Config = conf
}
