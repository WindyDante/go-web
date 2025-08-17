package config

import (
	"bytes"
	_ "embed"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	DBName    string `yaml:"dbname"`
	Charset   string `yaml:"charset"`
	ParseTime bool   `yaml:"parseTime"`
	Loc       string `yaml:"loc"`
}

type AuthConfig struct {
	Jwt struct {
		Expires  int    `yaml:"expires"`  // JWT的过期时间，单位为秒
		Issuer   string `yaml:"issuer"`   // JWT的发行者
		Audience string `yaml:"audience"` // JWT的受众
	} `yaml:"jwt"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Auth     AuthConfig     `yaml:"auth"`
	Redis    RedisConfig    `yaml:"redis"`
}

var AppConfig *Config

//go:embed config.yaml
var configFile []byte

// LoadConfig 加载配置文件
func LoadConfig() error {
	// 加载.env
	if err := godotenv.Load(); err != nil {
		return err
	}

	// 替换 YAML 中的 ${?} 占位符
	configFile = []byte(os.ExpandEnv(string(configFile)))

	// 读取YAML文件
	viper.SetConfigType("yaml")

	if err := viper.ReadConfig(bytes.NewReader(configFile)); err != nil {
		return err
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return err
	}

	return nil
}
