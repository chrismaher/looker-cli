package client

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
func readConfig() (*config, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	home := usr.HomeDir

	configFile := path.Join(home, ".looker/config.yaml")
	conf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	c := config{}
	err = yaml.Unmarshal(conf, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
