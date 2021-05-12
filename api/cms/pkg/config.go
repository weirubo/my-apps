package pkg

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	v *viper.Viper
}

type ServerConfig struct {
	HttpPort string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

type DatabaseConfig struct {
	Username string
	Password string
	Host string
	DBName string
	Charset string
	ParseTime bool
}

func NewViper() (*Config, error) {
	viper := viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../api/cms/configs/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("NewViper() error:", err)
		return nil, err
	}
	config := &Config{
		v: viper,
	}
	return config, nil
}

func (c *Config) ReadConfig(key string, value interface{}) error {
	err := c.v.UnmarshalKey(key, value)
	if err != nil {
		return err
	}
	return nil
}