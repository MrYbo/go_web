package dao

import (
	"login/app/database/mysql"
	"login/app/model/entity"
)

type User struct {}

func (u *User) FindOne(query interface{}, args ...interface{}) interface{} {
	var user entity.User
	if err := mysql.DB.Where(query, args...).Find(&user).Error; err != nil {
		return err
	}
	return user
}