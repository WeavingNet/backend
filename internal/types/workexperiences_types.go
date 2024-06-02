package types

import (
	"time"

	"github.com/zhufuyi/sponge/pkg/ggorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateWorkexperiencesRequest request params
type CreateWorkexperiencesRequest struct {
	UserID         int       `json:"userId" binding:""`         // 用户ID
	Company        string    `json:"company" binding:""`        // 公司
	Title          string    `json:"title" binding:""`          // 职位
	EmploymentType string    `json:"employmentType" binding:""` // 工作类型
	JobDescription string    `json:"jobDescription" binding:""` // 工作内容
	Location       string    `json:"location" binding:""`       // 地点
	StartDate      time.Time `json:"startDate" binding:""`      // 开始日期
	EndDate        time.Time `json:"endDate" binding:""`        // 结束日期
}

// UpdateWorkexperiencesByIDRequest request params
type UpdateWorkexperiencesByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id

	UserID         int       `json:"userId" binding:""`         // 用户ID
	Company        string    `json:"company" binding:""`        // 公司
	Title          string    `json:"title" binding:""`          // 职位
	EmploymentType string    `json:"employmentType" binding:""` // 工作类型
	JobDescription string    `json:"jobDescription" binding:""` // 工作内容
	Location       string    `json:"location" binding:""`       // 地点
	StartDate      time.Time `json:"startDate" binding:""`      // 开始日期
	EndDate        time.Time `json:"endDate" binding:""`        // 结束日期
}

// WorkexperiencesObjDetail detail
type WorkexperiencesObjDetail struct {
	ID string `json:"id"` // convert to string id

	UserID         int       `json:"userId"`         // 用户ID
	Company        string    `json:"company"`        // 公司
	Title          string    `json:"title"`          // 职位
	EmploymentType string    `json:"employmentType"` // 工作类型
	JobDescription string    `json:"jobDescription"` // 工作内容
	Location       string    `json:"location"`       // 地点
	StartDate      time.Time `json:"startDate"`      // 开始日期
	EndDate        time.Time `json:"endDate"`        // 结束日期
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// CreateWorkexperiencesRespond only for api docs
type CreateWorkexperiencesRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// UpdateWorkexperiencesByIDRespond only for api docs
type UpdateWorkexperiencesByIDRespond struct {
	Result
}

// GetWorkexperiencesByIDRespond only for api docs
type GetWorkexperiencesByIDRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Workexperiences WorkexperiencesObjDetail `json:"workexperiences"`
	} `json:"data"` // return data
}

// DeleteWorkexperiencesByIDRespond only for api docs
type DeleteWorkexperiencesByIDRespond struct {
	Result
}

// DeleteWorkexperiencessByIDsRequest request params
type DeleteWorkexperiencessByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// DeleteWorkexperiencessByIDsRespond only for api docs
type DeleteWorkexperiencessByIDsRespond struct {
	Result
}

// GetWorkexperiencesByConditionRequest request params
type GetWorkexperiencesByConditionRequest struct {
	query.Conditions
}

// GetWorkexperiencesByConditionRespond only for api docs
type GetWorkexperiencesByConditionRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Workexperiences WorkexperiencesObjDetail `json:"workexperiences"`
	} `json:"data"` // return data
}

// ListWorkexperiencessByIDsRequest request params
type ListWorkexperiencessByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// ListWorkexperiencessByIDsRespond only for api docs
type ListWorkexperiencessByIDsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Workexperiencess []WorkexperiencesObjDetail `json:"workexperiencess"`
	} `json:"data"` // return data
}

// ListWorkexperiencessRequest request params
type ListWorkexperiencessRequest struct {
	query.Params
}

// ListWorkexperiencessRespond only for api docs
type ListWorkexperiencessRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Workexperiencess []WorkexperiencesObjDetail `json:"workexperiencess"`
	} `json:"data"` // return data
}
