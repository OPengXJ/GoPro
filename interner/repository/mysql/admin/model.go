package admin

import "gorm.io/gorm"

type Admin struct{
	gorm.Model
	UserName string
	UserPass string
}