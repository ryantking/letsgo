package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config is the configuration for the different services.
type Config map[string]ServiceConfig

// LoadConfig loads the configurations stored in the given directory.
func LoadConfig(dir string) (Config, error) {
	path := fmt.Sprintf("%s/.letsgo.yaml", dir)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	cfg := make(Config)
	if err := yaml.NewDecoder(f).Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// ServiceConfig
type ServiceConfig struct {
	// The command to execute.
	Command string `yaml:"command"`

	// THe args to pass to the command.
	Args []string `yaml:"args"`

	// The directory to run the service from.
	Dir string `yaml:"dir"`

	// The environment to launch the program with.
	Env map[string]string `yaml:"env"`
}

// Environ returns the environment in the go expected format.
func (cfg *ServiceConfig) Environ() []string {
	environ := make([]string, 0, len(cfg.Env))
	for key, val := range cfg.Env {
		environ = append(environ, fmt.Sprintf("%s=%s", key, val))
	}

	return environ
}
