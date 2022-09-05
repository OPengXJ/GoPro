package main

import (
	"fmt"
	"time"

	"github.com/OPengXJ/GoPro/global"
	"github.com/OPengXJ/GoPro/pkg/setting"
)

func init(){
	setting,err:=setting.NewSetting()
	if err!=nil{
		fmt.Println(err)
	}
	if err:=setting.ReadSection("Database",&global.DatabaseSettings);err!=nil{
		fmt.Println(err)
	}

}

func main() {
	fmt.Println(global.DatabaseSettings)
	time.Sleep(10*time.Second)
}
