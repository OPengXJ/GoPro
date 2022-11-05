package admin

import (
	"github.com/OPengXJ/GoPro/interner/repository/mysql/admin"
)

type CreateAdminData struct{
	Username string	`form:"username"`
	Password string	`form:"password"`
}
func(s *Service)Create(data *CreateAdminData)error{
	model:=admin.NewModel()
	model.AdminName=data.Username
	model.AdminPass=data.Password
	err:=model.Create(s.db.Write)
	if err!=nil{
		return err
	}
	return nil
}