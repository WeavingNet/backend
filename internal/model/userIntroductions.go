package model

import (
	"github.com/zhufuyi/sponge/pkg/ggorm"
)

type UserIntroductions struct {
	ggorm.Model `gorm:"embedded"` // embed id and time

	UserID  int    `gorm:"column:user_id;type:int4;NOT NULL" json:"userId"`
	Title   string `gorm:"column:title;type:varchar(100);NOT NULL" json:"title"` // 介绍标题
	Content string `gorm:"column:content;type:text" json:"content"`              // 介绍内容
}
