package configs

import (
	"time"

	"github.com/spf13/viper"
)

var config = new(Config)

type Config struct {
	MySql struct {
		Bass struct {
			MaxIdleConnn    int `yaml:"maxidleconn"`
			MaxOpenConnn    int `yaml:"maxopenconn"`
			ConnMaxLifetime time.Duration `yaml:"connmaxlifetime"`
		} `yaml:"mysql.base"`
		Read struct {
			Host string `yaml:"Host"`
			User string `yaml:"User"`
			Pass string `yaml:"Pass"`
			Name string `yaml:"name"`
		} `yaml:"mysql.read"`
		Write struct {
			Host string `yaml:"Host"`
			User string `yaml:"User"`
			Pass string `yaml:"Pass"`
			Name string `yaml:"name"`
		} `yaml:"mysql.write"`
	} `yaml:"mysql"`
	JwtPass string `yaml:"jwtpass"`
	Redis struct{
		DBaddr string `yaml:"dbaddr"`
		DBpass string `yam;:"dbpass"`
		DB	int	`yaml:"db"`
		MaxRetries int `yaml:"maxretries"`
		PoolSize int `yaml:"poolsize"`
		MinIdleConns int `yaml:"minidleconns"`

	}`yaml:"redis"`
}

func init() {
	viper.AddConfigPath("configs/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}

func Get() Config {
	return *config
}
