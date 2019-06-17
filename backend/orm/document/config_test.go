package document

import (
	"testing"
)

func TestGetConfig(t *testing.T) {

	config, err := Config{}.GetConfig()
	if err != nil {
		t.Error(err)
	}

	t.Logf("config: %v \n", config)
}
