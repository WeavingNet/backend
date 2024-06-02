package model

import (
	"github.com/zhufuyi/sponge/pkg/ggorm"
)

type Skills struct {
	ggorm.Model `gorm:"embedded"` // embed id and time

	UserID           int    `gorm:"column:user_id;type:int4;NOT NULL" json:"userId"`                   // 用户ID
	SkillType        string `gorm:"column:skill_type;type:varchar(50);NOT NULL" json:"skillType"`      // 技能类型
	SkillName        string `gorm:"column:skill_name;type:varchar(50);NOT NULL" json:"skillName"`      // 技能名称
	ProficiencyLevel string `gorm:"column:proficiency_level;type:varchar(50)" json:"proficiencyLevel"` // 熟练程度
}
