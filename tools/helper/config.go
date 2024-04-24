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
		Autoremove  bool
	}
	Music struct {
		Vlc     string
		Source  string
		Enabled bool
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
	conf.Music.Vlc = os.ExpandEnv(conf.Music.Vlc)
	conf.Music.Source = os.ExpandEnv(conf.Music.Source)

	Config = conf
}

func StoreConfig() {
	viper.Set("fs.source", Config.FS.Source)
	viper.Set("fs.destination", Config.FS.Destination)
	viper.Set("fs.autoremove", Config.FS.Autoremove)
	viper.Set("music.vlc", Config.Music.Vlc)
	viper.Set("music.source", Config.Music.Source)
	viper.Set("music.enabled", Config.Music.Enabled)
	if err := viper.WriteConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
