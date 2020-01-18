package client

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os/user"
	"path"
)

type config struct {
	Path   string `yaml:"api_path"`
	ID     string `yaml:"client_id"`
	Secret string `yaml:"client_secret"`
}

// readConfig reads a config.yaml file stored in a .looker directory
// located in the user's home directory
func readConfig() *config {
	usr, err := user.Current()
	if err != nil {
		log.Println(err)
	}
	home := usr.HomeDir

	configFile := path.Join(home, ".looker/config.yaml")
	conf, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Println(err)
	}
	c := config{}
	err = yaml.Unmarshal(conf, &c)
	if err != nil {
		log.Println(err)
	}
	return &c
}
