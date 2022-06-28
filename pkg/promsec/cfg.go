package promsec

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

var (
	Config = viper.New()
)

func NewDefaultConfig() *viper.Viper {
	Config.SetDefault("version", "v1")
	Config.SetDefault("kind", "PromsecConfig")
	Config.SetDefault("server.host", "0.0.0.0")
	Config.SetDefault("server.port", 5001)
	Config.SetDefault("server.endpoint", "/metrics")
	Config.SetDefault("server.read_timeout", 5*time.Second)
	Config.SetDefault("server.read_header_timeout", 5*time.Second)
	Config.SetDefault("server.write_timeout", 5*time.Second)
	Config.SetDefault("server.idle_timeout", 10*time.Second)

	return Config
}

func NewConfigFromFile(f *os.File) (*viper.Viper, error) {
	if f == nil {
		return nil, fmt.Errorf("config file can not be nil")
	}

	if err := Config.ReadConfig(f); err != nil {
		return nil, err
	}

	return Config, nil
}
