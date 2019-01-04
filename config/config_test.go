package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetConfig tests config loading
func TestGetConfig(t *testing.T) {
	// Test env vars
	config, err := GetConfig()
	assert.Nil(t, err)
	assert.Equal(t, config.Hostname, ".*", "Wrong default hostname")
}
