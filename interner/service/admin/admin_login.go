package admin

import (
	"github.com/OPengXJ/GoPro/interner/repository/mysql/admin"
)

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginReponse struct {
	UserId   uint   `json:"userid"`
	UserName string `json:"username"`
}

func (s *Service) Login(req *LoginRequest) (*LoginReponse, error) {
	qb := admin.NewQueryBuilder()
	qb.WhereUsername(req.Username)
	qb.WherePassword(req.Password)
	admin, err := qb.First(s.db.Read)
	if err != nil {
		return nil, err
	}
	res := &LoginReponse{
		UserId:   admin.ID,
		UserName: admin.UserName,
	}
	return res, nil
}
