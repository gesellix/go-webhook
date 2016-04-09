package webhook

import (
	"encoding/json"
	"testing"
)

func TestReadConfig(t *testing.T) {
	config, err := ReadConfig("config-example.json")
	if err != nil {
		t.Errorf(err.Error())
	}
	configAsJson, err := json.Marshal(config)
	t.Logf("read config %v", string(configAsJson))
	if len(config.Handlers) != 2 {
		t.Errorf("wanted %d, got %d", 2, len(config.Handlers))
	}
}
