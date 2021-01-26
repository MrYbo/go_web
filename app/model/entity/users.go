package entity

import (
	"time"
)

type User struct {
	Id        int       `json:"id" gorm:"type:int" primary_key;auto_increment" description:"用户id"`
	Phone     string    `json:"phone"  gorm:"size:32;unique_index;not null" description:"手机号"`
	Name      string    `json:"name"      gorm:"size:32;not null" description:"用户名"`
	Password  string    `json:"password" gorm:"type:char(32);not null" description:"密码"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime" description:"创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime" description:"更新时间"`
	Invalid   string    `json:"invalid" gorm:"type:char(1);default:'N'" description:"是否有效"`
}
