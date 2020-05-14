package configuration

import (
	"log"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	var h interface{} = cfg.Server.Host
	s, ok := h.(string)

	if !ok {
		t.Errorf("unexpected error: %v", err)
	}

	t.Logf("unexpected error: %v", s)
}
