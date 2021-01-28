package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone    string `json:"phone"  gorm:"size:32;unique_index;not null" description:"手机号"`
	Name     string `json:"name"      gorm:"size:32;not null" description:"用户名"`
	Password string `json:"password" gorm:"type:char(32);not null" description:"密码"`
	Invalid  string `json:"invalid" gorm:"type:char(1);default:'N'" description:"是否有效"`
}
