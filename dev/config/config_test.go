package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigInvalidPath(t *testing.T) {
	err := LoadConfig("../invalid_path")
	assert.Error(t, err)
	assert.Empty(t, Configuration)
}

func TestLoadConfig(t *testing.T) {
	// Load configuration
	err := LoadConfig("../")
	assert.NoError(t, err, "Expected no error while loading configuration")

	// Assert server configuration
	serverConfig := Configuration.Server
	assert.Equal(t, 8080, serverConfig.Port)
}
