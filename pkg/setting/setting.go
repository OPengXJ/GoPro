package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct{
	vp *viper.Viper
}

func NewSetting()(*Setting,error){
	vp:=viper.New()
	vp.SetConfigType("yaml")
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	err:=vp.ReadInConfig()
	if err!=nil{
		return nil,err
	}else{
		vp.WatchConfig()
		vp.OnConfigChange(func(in fsnotify.Event) {
			fmt.Println("config has been changed !")
		})
		return &Setting{vp},nil
	}

}