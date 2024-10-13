package config

import "github.com/spf13/viper"

type DatabaseConfig struct {
	Dialect string `mapstructure:"dialect"`
	DSN     string `mapstructure:"dsn"`
}

type LoggerConfig struct {
	LogLevel      string `mapstructure:"log_level"`
	SlowThreshold int    `mapstructure:"slow_threshold"`
}

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Logger   LoggerConfig   `mapstructure:"logger"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // 配置文件名（不带扩展名）
	viper.SetConfigType("toml")   // 配置文件类型
	viper.AddConfigPath("config") // 配置文件路径

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	println("aaa", &config)

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
