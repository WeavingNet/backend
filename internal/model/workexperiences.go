package model

import (
	"github.com/zhufuyi/sponge/pkg/ggorm"
	"time"
)

type Workexperiences struct {
	ggorm.Model `gorm:"embedded"` // embed id and time

	UserID         int       `gorm:"column:user_id;type:int4;NOT NULL" json:"userId"`               // 用户ID
	Company        string    `gorm:"column:company;type:varchar(100);NOT NULL" json:"company"`      // 公司
	Title          string    `gorm:"column:title;type:varchar(50)" json:"title"`                    // 职位
	EmploymentType string    `gorm:"column:employment_type;type:varchar(50)" json:"employmentType"` // 工作类型
	JobDescription string    `gorm:"column:job_description;type:text" json:"jobDescription"`        // 工作内容
	Location       string    `gorm:"column:location;type:varchar(100)" json:"location"`             // 地点
	StartDate      time.Time `gorm:"column:start_date;type:date" json:"startDate"`                  // 开始日期
	EndDate        time.Time `gorm:"column:end_date;type:date" json:"endDate"`                      // 结束日期
}
