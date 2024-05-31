package model

import (
	"github.com/zhufuyi/sponge/pkg/ggorm"
)

type Users struct {
	ggorm.Model `gorm:"embedded"` // embed id and time

	FirstName         string `gorm:"column:first_name;type:varchar(50);NOT NULL" json:"firstName"`          // 名字
	LastName          string `gorm:"column:last_name;type:varchar(50);NOT NULL" json:"lastName"`            // 姓氏
	ProfilePictureUrl string `gorm:"column:profile_picture_url;type:varchar(255)" json:"profilePictureUrl"` // 头像URL
	About             string `gorm:"column:about;type:text" json:"about"`                                   // 个人简介
}
