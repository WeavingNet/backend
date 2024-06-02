package types

import (
	"time"

	"github.com/zhufuyi/sponge/pkg/ggorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateProjectsRequest request params
type CreateProjectsRequest struct {
	UserID      int    `json:"userId" binding:""`      // 用户ID
	ProjectName string `json:"projectName" binding:""` // 项目名称
	Role        string `json:"role" binding:""`        // 所担任角色
	Description string `json:"description" binding:""` // 项目介绍/成就
}

// UpdateProjectsByIDRequest request params
type UpdateProjectsByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id

	UserID      int    `json:"userId" binding:""`      // 用户ID
	ProjectName string `json:"projectName" binding:""` // 项目名称
	Role        string `json:"role" binding:""`        // 所担任角色
	Description string `json:"description" binding:""` // 项目介绍/成就
}

// ProjectsObjDetail detail
type ProjectsObjDetail struct {
	ID string `json:"id"` // convert to string id

	UserID      int       `json:"userId"`      // 用户ID
	ProjectName string    `json:"projectName"` // 项目名称
	Role        string    `json:"role"`        // 所担任角色
	Description string    `json:"description"` // 项目介绍/成就
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// CreateProjectsRespond only for api docs
type CreateProjectsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// UpdateProjectsByIDRespond only for api docs
type UpdateProjectsByIDRespond struct {
	Result
}

// GetProjectsByIDRespond only for api docs
type GetProjectsByIDRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Projects ProjectsObjDetail `json:"projects"`
	} `json:"data"` // return data
}

// DeleteProjectsByIDRespond only for api docs
type DeleteProjectsByIDRespond struct {
	Result
}

// DeleteProjectssByIDsRequest request params
type DeleteProjectssByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// DeleteProjectssByIDsRespond only for api docs
type DeleteProjectssByIDsRespond struct {
	Result
}

// GetProjectsByConditionRequest request params
type GetProjectsByConditionRequest struct {
	query.Conditions
}

// GetProjectsByConditionRespond only for api docs
type GetProjectsByConditionRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Projects ProjectsObjDetail `json:"projects"`
	} `json:"data"` // return data
}

// ListProjectssByIDsRequest request params
type ListProjectssByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// ListProjectssByIDsRespond only for api docs
type ListProjectssByIDsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Projectss []ProjectsObjDetail `json:"projectss"`
	} `json:"data"` // return data
}

// ListProjectssRequest request params
type ListProjectssRequest struct {
	query.Params
}

// ListProjectssRespond only for api docs
type ListProjectssRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Projectss []ProjectsObjDetail `json:"projectss"`
	} `json:"data"` // return data
}
