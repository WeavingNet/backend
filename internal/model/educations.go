package model

import (
	"github.com/zhufuyi/sponge/pkg/ggorm"
	"time"
)

type Educations struct {
	ggorm.Model `gorm:"embedded"` // embed id and time

	UserID       int       `gorm:"column:user_id;type:int4;NOT NULL" json:"userId"`            // 用户ID
	School       string    `gorm:"column:school;type:varchar(100);NOT NULL" json:"school"`     // 学校
	Degree       string    `gorm:"column:degree;type:varchar(50)" json:"degree"`               // 学位
	FieldOfStudy string    `gorm:"column:field_of_study;type:varchar(50)" json:"fieldOfStudy"` // 专业
	StartDate    time.Time `gorm:"column:start_date;type:date" json:"startDate"`               // 开始日期
	EndDate      time.Time `gorm:"column:end_date;type:date" json:"endDate"`                   // 结束日期
	Gpa          string    `gorm:"column:gpa;type:numeric" json:"gpa"`                         // 平均成绩
	Activities   string    `gorm:"column:activities;type:text" json:"activities"`              // 活动/社团
}
