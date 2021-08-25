package env

import (
	"time"

	"github.com/spf13/viper"
)

//Config with all environmental variables
type Config struct {
	ServerAddress        string        `mapstructure:"server_address"`
	DbUser               string        `mapstructure:"db_user"`
	DbPassword           string        `mapstructure:"db_password"`
	DbHost               string        `mapstructure:"db_host"`
	DbPort               int           `mapstructure:"db_port"`
	DbName               string        `mapstructure:"db_name"`
	MaxConnectionPool    int           `mapstructure:"max_connection_pool"`
	MaxConnectionTimeout time.Duration `mapstructure:"max_connection_timeout"`
	AuthUser             string        `mapstructure:"auth_user"`
	AuthPassword         string        `mapstructure:"auth_password"`
}

func (c *Config) setDefaults() {
	defaults := map[string]interface{}{
		"server_address":         ":8080",
		"db_host":                "localhost",
		"db_port":                5432,
		"db_name":                "seen",
		"max_connection_pool":    3,
		"max_connection_timeout": time.Minute,
	}
	for key, value := range defaults {
		viper.SetDefault(key, value)
	}
}

func (c *Config) setFromEnvironment() error {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigFile("env")
	viper.SetConfigName("secrets")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	c.ServerAddress = viper.GetString("server_address")
	c.DbUser = viper.GetString("db_user")
	c.DbPassword = viper.GetString("db_password")
	c.DbHost = viper.GetString("db_host")
	c.DbPort = viper.GetInt("db_port")
	c.DbName = viper.GetString("db_name")
	c.MaxConnectionPool = viper.GetInt("max_connection_pool")
	c.MaxConnectionTimeout = viper.GetDuration("max_connection_timeout")
	c.AuthUser = viper.GetString("auth_user")
	c.AuthPassword = viper.GetString("auth_password")

	if err := viper.Unmarshal(&c); err != nil {
		return err
	}

	return nil
}

// NewConfig initializes environment
func NewConfig() (*Config, error) {
	config := &Config{}
	config.setDefaults()
	if err := config.setFromEnvironment(); err != nil {
		return nil, err
	}
	return config, nil
}
