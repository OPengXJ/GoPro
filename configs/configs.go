package configs

import "github.com/spf13/viper"

var config = new(Config)

type Config struct{
	MySql struct{
		Read struct{
			Host string `yaml:"Host"`
			User string `yaml:"User"`
			Pass string	`yaml:"Pass"`
			Name string	`yaml:"name"`
		}`yaml:"mysql.read"`
		Write struct{
			Host string `yaml:"Host"`
			User string `yaml:"User"`
			Pass string	`yaml:"Pass"`
			Name string	`yaml:"name"`
		}`yaml:"mysql.write"`
	}`yaml:"mysql"`
}
func init(){
	viper.AddConfigPath("configs/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	err:=viper.ReadInConfig()
	if err !=nil{
		panic(err)
	}
	err=viper.Unmarshal(&config)
	if err!=nil{
		panic(err)
	}
}

func Get() Config {
	return *config
}