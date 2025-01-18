package conf

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppName  string       `yaml:"appname"`
	Server   ServerConfig `yaml:"server"`
	Database Database     `yaml:"database"`
	Redis    RedisConfig  `yaml:"redis"`
	JWT      JWTConfig    `yaml:"jwt"`
}

type Database struct {
	Host     string   `yaml:"host"`
	Port     int      `yaml:"port"`
	User     string   `yaml:"user"`
	Password string   `yaml:"password"`
	DBName   string   `yaml:"dbname"`
	Scripts  []string `yaml:"scripts"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type ServerConfig struct {
	Port         int  `yaml:"port"`
	ReadTimeout  int  `yaml:"readtimeout"`
	WriteTimeout int  `yaml:"writetimeout"`
	RateLimit    bool `yaml:"ratelimit"`
}

type JWTConfig struct {
	Secret string `yaml:"secret"`
}

func LoadConfig(configPath string) (*Config, error) {
	var config Config

	// Normalize the config path for cross-platform compatibility
	absolutePath, err := filepath.Abs(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Set up Viper
	viper.SetConfigFile(filepath.Join(absolutePath, "config.yaml")) // Explicitly set the file path

	// Automatically bind environment variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Warning: no configuration file found (%v). Using environment variables.\n", err)
	}

	// Unmarshal the configuration into the struct
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &config, nil
}
