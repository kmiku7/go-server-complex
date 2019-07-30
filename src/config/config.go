package config

import (
	"github.com/BurntSushi/toml"
)

type CoreConfig struct {
	ListenAddress string `toml:"listen_address"`
}

type LogConfig struct {
	LogDir string `toml:"log_dir"`
}

type Config struct {
	Core *CoreConfig `toml:"core"`
	Log  *LogConfig  `toml:"log"`
}

var (
	globalConfig *Config
)

func LoadConfig(filepath string) (*Config, error) {

	config := &Config{}
	if _, err := toml.DecodeFile(filepath, config); err != nil {
		return nil, err
	}
	return config, nil
}

func InitGlobalConfig(filepath string) error {
	config, err := LoadConfig(filepath)
	if err != nil {
		return err
	}
	globalConfig = config
	return nil
}

func GetGlobalConfig() *Config {
	return globalConfig
}
