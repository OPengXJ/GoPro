package mysql

import (
	"errors"
	"fmt"
	"time"

	"github.com/OPengXJ/GoPro/configs"
	"github.com/OPengXJ/GoPro/interner/repository/mysql/admin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repo struct{
	Write *gorm.DB
	Read *gorm.DB
}

var repo =new(Repo)

func New() (*Repo,error){
	mysql_config:=configs.Get().MySql
	dbw,err:=dbConnect(mysql_config.Write.Host,mysql_config.Write.User,mysql_config.Write.Pass,mysql_config.Write.Name)
	if err!=nil{
		return nil,errors.New("dbwrite connection failed !")
	}
	dbr,err:=dbConnect(mysql_config.Read.Host,mysql_config.Read.User,mysql_config.Read.Pass,mysql_config.Read.Name)
	if err!=nil{
		return nil,errors.New("dbread connection failde !")
	}
	return &Repo{
		Write: dbw,
		Read: dbr,
	},nil
}

func dbConnect(host string,user string,pass string,name string)(*gorm.DB,error){
	dsn:=fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",user,pass,host,name)
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		return nil,fmt.Errorf("DB connection failed ! DBname : %s",name)
	}
	err=db.AutoMigrate(&admin.Admin{})
	if err!=nil{
		fmt.Println("建表error",err)
	}
	sqlDb,err:=db.DB()
	if err!=nil{
		fmt.Println(err)
	}
	mysql_base:=configs.Get().MySql.Bass
	sqlDb.SetConnMaxLifetime(mysql_base.ConnMaxLifetime*time.Second)
	sqlDb.SetMaxIdleConns(mysql_base.MaxIdleConnn)
	sqlDb.SetMaxOpenConns(mysql_base.MaxOpenConnn)
	return db,nil
}
func init(){
	var err error
	repo,err=New()
	if err!=nil{
		panic(err)
	}
}
func GetMysqlRepo()*Repo{
	return repo
}
