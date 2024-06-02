package types

import (
	"time"

	"github.com/zhufuyi/sponge/pkg/ggorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateEducationsRequest request params
type CreateEducationsRequest struct {
	UserID       int       `json:"userId" binding:""`       // 用户ID
	School       string    `json:"school" binding:""`       // 学校
	Degree       string    `json:"degree" binding:""`       // 学位
	FieldOfStudy string    `json:"fieldOfStudy" binding:""` // 专业
	StartDate    time.Time `json:"startDate" binding:""`    // 开始日期
	EndDate      time.Time `json:"endDate" binding:""`      // 结束日期
	Gpa          string    `json:"gpa" binding:""`          // 平均成绩
	Activities   string    `json:"activities" binding:""`   // 活动/社团
}

// UpdateEducationsByIDRequest request params
type UpdateEducationsByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id

	UserID       int       `json:"userId" binding:""`       // 用户ID
	School       string    `json:"school" binding:""`       // 学校
	Degree       string    `json:"degree" binding:""`       // 学位
	FieldOfStudy string    `json:"fieldOfStudy" binding:""` // 专业
	StartDate    time.Time `json:"startDate" binding:""`    // 开始日期
	EndDate      time.Time `json:"endDate" binding:""`      // 结束日期
	Gpa          string    `json:"gpa" binding:""`          // 平均成绩
	Activities   string    `json:"activities" binding:""`   // 活动/社团
}

// EducationsObjDetail detail
type EducationsObjDetail struct {
	ID string `json:"id"` // convert to string id

	UserID       int       `json:"userId"`       // 用户ID
	School       string    `json:"school"`       // 学校
	Degree       string    `json:"degree"`       // 学位
	FieldOfStudy string    `json:"fieldOfStudy"` // 专业
	StartDate    time.Time `json:"startDate"`    // 开始日期
	EndDate      time.Time `json:"endDate"`      // 结束日期
	Gpa          string    `json:"gpa"`          // 平均成绩
	Activities   string    `json:"activities"`   // 活动/社团
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// CreateEducationsRespond only for api docs
type CreateEducationsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// UpdateEducationsByIDRespond only for api docs
type UpdateEducationsByIDRespond struct {
	Result
}

// GetEducationsByIDRespond only for api docs
type GetEducationsByIDRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Educations EducationsObjDetail `json:"educations"`
	} `json:"data"` // return data
}

// DeleteEducationsByIDRespond only for api docs
type DeleteEducationsByIDRespond struct {
	Result
}

// DeleteEducationssByIDsRequest request params
type DeleteEducationssByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// DeleteEducationssByIDsRespond only for api docs
type DeleteEducationssByIDsRespond struct {
	Result
}

// GetEducationsByConditionRequest request params
type GetEducationsByConditionRequest struct {
	query.Conditions
}

// GetEducationsByConditionRespond only for api docs
type GetEducationsByConditionRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Educations EducationsObjDetail `json:"educations"`
	} `json:"data"` // return data
}

// ListEducationssByIDsRequest request params
type ListEducationssByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// ListEducationssByIDsRespond only for api docs
type ListEducationssByIDsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Educationss []EducationsObjDetail `json:"educationss"`
	} `json:"data"` // return data
}

// ListEducationssRequest request params
type ListEducationssRequest struct {
	query.Params
}

// ListEducationssRespond only for api docs
type ListEducationssRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Educationss []EducationsObjDetail `json:"educationss"`
	} `json:"data"` // return data
}
