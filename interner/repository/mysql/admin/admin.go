package admin

import "gorm.io/gorm"

func NewModel() *Admin{
	return new(Admin)
}

func(t *Admin)Create(db *gorm.DB)error{
	result:=db.Create(t)
	if result.Error!=nil{
		return result.Error
	}
	return nil
}