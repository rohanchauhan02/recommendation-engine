package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

// ImmutableConfig is an interface represent methods in config
type ImmutableConfig interface {
	GetPort() string
	GetDatabase() Database
}

type Config struct {
	Port     string   `mapstructure:"PORT"`
	Database Database `mapstructure:"DATABASE"`
}

type Database struct {
	Mysql MysqlConf `mapstructure:"MYSQL"`
	Neo4j Neo4jConf `mapstructure:"NEO4J"`
}

type MysqlConf struct {
	Port     string `mapstructure:"PORT"`
	Host     string `mapstructure:"HOST"`
	User     string `mapstructure:"USER"`
	Name     string `mapstructure:"NAME"`
	Password string `mapstructure:"PASS"`
}

type Neo4jConf struct {
	Uri      string `mapstructure:"URI"`
	Name     string `mapstructure:"NAME"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASS"`
}

func (i Config) GetPort() string {
	return i.Port
}

func (i Config) GetDatabase() Database {
	return i.Database
}

var (
	imOnce    sync.Once
	immutable *Config
)

// NewImmutableConfig is a factory that return of its config implementation
func NewImmutableConfig() ImmutableConfig {
	imOnce.Do(func() {
		v := viper.New()
		appEnv, exists := os.LookupEnv("APP_ENV")
		configName := "app.config.local"
		if exists {
			if appEnv == "production" {
				configName = "app.config.prod"
			} else {
				configName = "app.config.dev"
			}
		}
		v.SetConfigName("configs/" + configName)
		v.AddConfigPath(".")
		v.SetEnvPrefix("GO_RE")
		v.AutomaticEnv()
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := v.ReadInConfig(); err != nil {
			fmt.Println(500, err, "[CONFIG][missing] Failed to read app.config.* file", "failed")
		}
		v.Unmarshal(&immutable)

	})
	return immutable
}
