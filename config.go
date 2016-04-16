package webhook

import (
	"encoding/json"
	"io/ioutil"
)

//go:generate jsonenums -type=MessageType
type MessageType int

const (
	DockerHub MessageType = iota
	Drone
)

type Tls struct {
	Key  string `json:"key,omitempty"`
	Cert string `json:"cert,omitempty"`
}

type Action struct {
	Command string `json:"command,omitempty"`
}

type Handler struct {
	Path    string      `json:"path"`
	Type    MessageType `json:"type"`
	ApiKey  string      `json:"apikey,omitempty"`
	Actions []Action    `json:"actions,omitempty"`
}

type Config struct {
	Tls      Tls       `json:"tls,omitempty"`
	Handlers []Handler `json:"handlers"`
}

func ReadConfig(file string) (*Config, error) {
	config := &Config{}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return config, json.Unmarshal(bytes, &config)
}
