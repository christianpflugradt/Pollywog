package sys

import (
	"gopkg.in/yaml.v2"
	_ "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"pollywog/util"
)

type Config struct {
	Client struct {
		BaseUrl string `yaml:"baseurl"`
	}
	Server struct {
		Port string `yaml:"port"`
		Admintoken string `yaml:"admintoken"`
		Admintokens []Admintoken
	}
	Database struct {
		Driver string `yaml:"driver"`
		DataSourceName string `yaml:"dataSourceName"`
	}
	Smtp struct {
		Identity string `yaml:"identity"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
}

type Admintoken struct {
	User string `yaml:"user"`
	Token string `yaml:"token"`
	Whitelist []string `yaml:"whitelist"`
}

func (config *Config) Get() *Config {
	yml, err := ioutil.ReadFile(os.Args[1])
	util.HandleError(util.ErrorLogEvent{ Function: "config.Get", Error: err })
	err = yaml.Unmarshal(yml, &config)
	util.HandleError(util.ErrorLogEvent{ Function: "config.Get", Error: err })
	return config
}
