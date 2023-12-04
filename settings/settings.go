package settings

import (
	_ "embed"

	"github.com/go-yaml/yaml"
)

//go:embed settings.yml
var settingFile []byte

type Settings struct{
	DB DatabaseConfig    `yaml:"database"`
	Port int			 `yaml:"port"`
}

type DatabaseConfig struct{
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Name string `yaml:"name"`
}

func NewSettings() (*Settings,error){
	settings := &Settings{}
	err := yaml.Unmarshal(settingFile,&settings)
	if err != nil {
		return nil, err
	}
	return settings,nil
}