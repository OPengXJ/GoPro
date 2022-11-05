package admin

import "gorm.io/gorm"

type Admin struct{
	gorm.Model
	AdminName string
	AdminPass string
}