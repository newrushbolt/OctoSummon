package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

// MainConfig is root config
type MainConfig struct {
	Hostname string `env:"OCTOSUMMON_HOSTNAME" envDefault:".*"`
}

// GetConfig is a init func, returning root config
func GetConfig() (MainConfig, error) {
	config := MainConfig{}

	var localErr error
	err := env.Parse(&config)
	if err != nil {
		localErr = fmt.Errorf("Cannot parse ENV:\n%v", err)
	}
	return config, localErr

}
