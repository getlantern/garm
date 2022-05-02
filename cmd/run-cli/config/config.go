package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"

	runnerErrors "runner-manager/errors"
)

const (
	DefaultAppFolder      = "run-cli"
	DefaultConfigFileName = "config.toml"
)

func getConfigFilePath() (string, error) {
	configDir, err := getHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "fetching home folder")
	}

	if err := ensureHomeDir(configDir); err != nil {
		return "", errors.Wrap(err, "ensuring config dir")
	}

	cfgFile := filepath.Join(configDir, DefaultConfigFileName)
	return cfgFile, nil
}

func LoadConfig() (*Config, error) {
	cfgFile, err := getConfigFilePath()
	if err != nil {
		return nil, errors.Wrap(err, "fetching config")
	}

	var config Config
	if _, err := toml.DecodeFile(cfgFile, &config); err != nil {
		return nil, errors.Wrap(err, "decoding toml")
	}

	return &config, nil
}

type Config struct {
	Managers      []Manager `toml:"manager"`
	ActiveManager string    `toml:"active_manager"`
}

func (c *Config) HasManager(mgr string) bool {
	if mgr == "" {
		return false
	}
	for _, val := range c.Managers {
		if val.Name == mgr {
			return true
		}
	}
	return false
}

func (c *Config) GetActiveConfig() (Manager, error) {
	if c.ActiveManager == "" {
		return Manager{}, runnerErrors.ErrNotFound
	}

	for _, val := range c.Managers {
		if val.Name == c.ActiveManager {
			return val, nil
		}
	}
	return Manager{}, runnerErrors.ErrNotFound
}

func (c *Config) SaveConfig() error {
	cfgFile, err := getConfigFilePath()
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return errors.Wrap(err, "getting config")
		}
	}
	cfgHandle, err := os.Create(cfgFile)
	if err != nil {
		errors.Wrap(err, "getting file handle")
	}

	encoder := toml.NewEncoder(cfgHandle)
	if err := encoder.Encode(c); err != nil {
		return errors.Wrap(err, "saving config")
	}

	return nil
}

type Manager struct {
	Name    string `toml:"name"`
	BaseURL string `toml:"base_url"`
	Token   string `toml:"bearer_token"`
}
