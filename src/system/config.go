package sys

import (
	"fmt"
	"gopkg.in/yaml.v2"
	_ "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Client struct {
		BaseUrl string `yaml:"baseurl"`
	}
	Server struct {
		Port string `yaml:"port"`
	}
	Smtp struct {
		Identity string `yaml:"identity"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
}

func (config *Config) Get() *Config {
	yml, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Print(err)
	}
	err = yaml.Unmarshal(yml, &config)
	if err != nil {
		fmt.Print(err)
	}
	return config
}
