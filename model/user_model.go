package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 构造 用户 model
type User struct {
	gorm.Model
	Username    string    `json:"username" gorm:"column:'username'"`
	Password    string    `json:"password" gorm:"column:'password'"`
	Gender      bool      `json:"gender" gorm:"column:'gender';default:false"`
	IsDelete    bool      `json:"is_delete" gorm:"column:'Is_delete';default:false"`
	IsSuper     bool      `json:"is_super" gorm:"column:'is_super';default:false"`
	DeletedTime time.Time `json:"deleted_time" gorm:"column:'deleted_time'"`
}
