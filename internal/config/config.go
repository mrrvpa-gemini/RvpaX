package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ScanConfig struct {
	URL       string
	ListFile  string
	DeepScan  bool
	JSON      bool
	HTML      bool
	Output    string
	Threads   int
	Timeout   int
	Proxy     string
	UserAgent string
}

type HTTPConfig struct {
	Timeout   int    `yaml:"timeout"`
	Retries   int    `yaml:"retries"`
	UserAgent string `yaml:"user_agent"`
}

type ScanSettings struct {
	Threads  int  `yaml:"threads"`
	DeepScan bool `yaml:"deep_scan"`
	Timeout  int  `yaml:"timeout"`
}

type OutputSettings struct {
	JSON   bool   `yaml:"json"`
	HTML   bool   `yaml:"html"`
	Format string `yaml:"format"`
}

type Config struct {
	HTTP   HTTPConfig      `yaml:"http"`
	Scan   ScanSettings    `yaml:"scan"`
	Output OutputSettings `yaml:"output"`
}

func LoadConfig() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(home, ".rvpa")
	configFile := filepath.Join(configDir, "config.yaml")

	// Create config directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}

	// Create default config if it doesn't exist
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		defaultCfg := &Config{
			HTTP: HTTPConfig{
				Timeout:   30,
				Retries:   3,
				UserAgent: "RVPA/1.0",
			},
			Scan: ScanSettings{
				Threads:  20,
				DeepScan: false,
				Timeout:  30,
			},
			Output: OutputSettings{
				JSON:   false,
				HTML:   false,
				Format: "text",
			},
		}
		data, _ := yaml.Marshal(defaultCfg)
		os.WriteFile(configFile, data, 0644)
		return defaultCfg, nil
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
