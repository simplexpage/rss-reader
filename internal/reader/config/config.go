package config

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	ListenHttp struct {
		Port string `yaml:"port" env-default:"8082"`
	} `yaml:"listen_http"`
	Rabbit struct {
		Host     string `yaml:"host" env-default:"rabbitmq"`
		Port     string `yaml:"port" env-default:"5672"`
		User     string `yaml:"user" env-default:"guest"`
		Password string `yaml:"password" env-default:"guest"`
	} `yaml:"rabbit"`
}

var instance *Config
var once sync.Once

func GetConfig(logger log.Logger) *Config {
	once.Do(func() {
		level.Info(logger).Log("msg", "Loading config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			level.Info(logger).Log("msg", help)
			level.Error(logger).Log("err", err)
		}
		level.Info(logger).Log("msg", "Config loaded")
	})
	return instance
}
