package config

import (
	"fmt"
	// "io/ioutil"
	// "regexp"

	"github.com/caarlos0/env"
	// "github.com/newrushbolt/OctoSummon/logger"
	// "gopkg.in/yaml.v2"
)

// MainConfig is root config
type MainConfig struct {
	Hostname string `env:"OCTOSUMMON_HOSTNAME" envDefault:"{.*}"`
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

	// yamlFile, err := ioutil.ReadFile(configPath)
	// if err != nil {
	// 	localErr := fmt.Errorf("Cannot read <%s>:\n%v ", configPath, err)
	// 	return config, localErr
	// }
	// logger.Logger.Debugf("Got YAML config:\n%s\n", yamlFile)

	// err = yaml.Unmarshal(yamlFile, &config)
	// if err != nil {
	// 	localErr := fmt.Errorf("Cannot parse YAML:\n%s\n%v", yamlFile, err)
	// 	return config, localErr
	// }

}
