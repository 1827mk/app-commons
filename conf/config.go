package conf

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Auth configuration
type AuthConfig struct {
	TokenExpiry    int `yaml:"tokenExpiry"`    // Token expiry in hours
	RefreshExpiry  int `yaml:"refreshExpiry"`  // Refresh token expiry in days
	MaxDevices     int `yaml:"maxDevices"`     // Maximum devices per user
	SessionTimeout int `yaml:"sessionTimeout"` // Session timeout in minutes
	RateLimit      struct {
		LoginAttempts int `yaml:"loginAttempts"` // Max login attempts
		WindowMinutes int `yaml:"windowMinutes"` // Time window for rate limiting
	} `yaml:"rateLimit"`
}

// Enhanced JWT configuration
type JWTConfig struct {
	Secret        string `yaml:"secret"`
	AccessExpiry  int    `yaml:"accessExpiry"`  // Access token expiry in minutes
	RefreshExpiry int    `yaml:"refreshExpiry"` // Refresh token expiry in days
	Issuer        string `yaml:"issuer"`
	Audience      string `yaml:"audience"`
}

type Config struct {
	AppName  string       `yaml:"appname"`
	Server   ServerConfig `yaml:"server"`
	Database Database     `yaml:"database"`
	Redis    RedisConfig  `yaml:"redis"`
	JWT      JWTConfig    `yaml:"jwt"`
	Auth     AuthConfig   `yaml:"auth"`
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
