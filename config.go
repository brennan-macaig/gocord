package gocord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Command         string   `json:"command"`
	IgnoredUsers    []string `json:"ignoredUsers"`
	IgnoredChannels []string `json:"ignoredChannels"`
	DatabasePath    string   `json:"databasePath"`
}

type Secrets struct {
	Token string `json:"token"`
}

func GetConfig(path string) (Config, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("could not read config file - %w", err)
	}
	var conf Config
	err = json.Unmarshal(dat, &conf)
	if err != nil {
		return Config{}, fmt.Errorf("could not unmarshal config file - %w", err)
	}
	if len(conf.Command) < 1 {
		return Config{}, fmt.Errorf("command must be at least 1 character long")
	}
	return conf, nil
}

func GetSecret(path string) (Secrets, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return Secrets{}, fmt.Errorf("could not read secrets file - %w", err)
	}
	var sec Secrets
	err = json.Unmarshal(dat, &sec)
	if err != nil {
		return Secrets{}, fmt.Errorf("could not unmarshal secrets file - %w", err)
	}
	if len(sec.Token) < 1 {
		return Secrets{}, fmt.Errorf("must have some discord token in secrets file")
	}
	return sec, nil
}

func (o Config) userIgnored(id string) bool {
	for _, val := range o.IgnoredUsers {
		if id == val {
			return true
		}
	}
	return false
}

func (o Config) channelIgnored(id string) bool {
	for _, val := range o.IgnoredChannels {
		if id == val {
			return true
		}
	}
	return false
}
