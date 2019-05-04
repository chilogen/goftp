package config

import (
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
	"os"
)

type cfg struct {
	Mysql struct {
		Infodb string `yaml:"infodb"`
	} `yaml:"Mysql"`

	FlowCount struct {
		MaxCon int `yamk:"maxcon"`
	} `yaml:"FlowCount"`
}

var (
	configInstance cfg
	LoadPath       = "./conf/config.yml"
	logger         = logrus.New()
)

func init() {
	err := Load(&configInstance)
	if err != nil {
		return
	}
	logger.Info("go init complete config")

}

func Load(config interface{}) error {
	env := os.Getenv("ENV")
	if env == "" {
		env = "staging"
	}
	return configor.New(&configor.Config{Environment: env, Verbose: false}).Load(config, LoadPath)
}

func GetConfig() *cfg {
	return &configInstance
}
