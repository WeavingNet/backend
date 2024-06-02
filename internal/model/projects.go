package model

import (
	"github.com/zhufuyi/sponge/pkg/ggorm"
)

type Projects struct {
	ggorm.Model `gorm:"embedded"` // embed id and time

	UserID      int    `gorm:"column:user_id;type:int4;NOT NULL" json:"userId"`                   // 用户ID
	ProjectName string `gorm:"column:project_name;type:varchar(100);NOT NULL" json:"projectName"` // 项目名称
	Role        string `gorm:"column:role;type:varchar(50)" json:"role"`                          // 所担任角色
	Description string `gorm:"column:description;type:text" json:"description"`                   // 项目介绍/成就
}
