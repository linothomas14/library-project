package config

import (
	"github.com/spf13/viper"
)

const ConfigName = "config"
const ConfigType = "yaml"

var Configuration Config

type Config struct {
	DB     DBConfig     `mapstructure:"db"`
	Server Serverconfig `mapstructure:"server"`
}

type DBConfig struct {
	DBHost          string `mapstructure:"dbhost"`
	DBName          string `mapstructure:"dbname"`
	DBPort          int    `mapstructure:"dbport"`
	DBUser          string `mapstructure:"dbuser"`
	DBPass          string `mapstructure:"dbpass"`
	AuthSource      string `mapstructure:"auth_source"`
	LibraryDatabase string `mapstructure:"library_db"`
	BookCollection  string `mapstructure:"book_collection"`
}
type Serverconfig struct {
	Port int `mapstructure:"port"`
}

func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}
	Configuration = config
	return nil
}
