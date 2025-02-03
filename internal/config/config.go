package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Http_server struct {
	Addr string	`yaml:"address"`
}

type Config struct {
	Env         string `yaml:"env"`
	Http_server `yaml:"http_server"`
}

func MustLoad() *Config{
	var configPath string = os.Getenv("CONFIG")

	if configPath == ""{
		log.Fatal("Error cannot get config path")
	}
	var cfg Config
	if err:=cleanenv.ReadConfig(configPath,&cfg); err!=nil{
		log.Fatal(err)
	}
	return &cfg
}
