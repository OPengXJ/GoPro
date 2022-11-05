package admin

import "github.com/OPengXJ/GoPro/interner/repository/mysql"


type Service struct{
	db mysql.Repo
}

func New(db mysql.Repo)*Service{
	return &Service{
		db: db,
	}
}