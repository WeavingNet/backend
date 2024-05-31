package types

import (
	"time"

	"github.com/zhufuyi/sponge/pkg/ggorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateUsersRequest request params
type CreateUsersRequest struct {
	FirstName         string `json:"firstName" binding:""`         // 名字
	LastName          string `json:"lastName" binding:""`          // 姓氏
	ProfilePictureUrl string `json:"profilePictureUrl" binding:""` // 头像URL
	About             string `json:"about" binding:""`             // 个人简介
}

// UpdateUsersByIDRequest request params
type UpdateUsersByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id

	FirstName         string `json:"firstName" binding:""`         // 名字
	LastName          string `json:"lastName" binding:""`          // 姓氏
	ProfilePictureUrl string `json:"profilePictureUrl" binding:""` // 头像URL
	About             string `json:"about" binding:""`             // 个人简介
}

// UsersObjDetail detail
type UsersObjDetail struct {
	ID string `json:"id"` // convert to string id

	FirstName         string    `json:"firstName"`         // 名字
	LastName          string    `json:"lastName"`          // 姓氏
	ProfilePictureUrl string    `json:"profilePictureUrl"` // 头像URL
	About             string    `json:"about"`             // 个人简介
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

// CreateUsersRespond only for api docs
type CreateUsersRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// UpdateUsersByIDRespond only for api docs
type UpdateUsersByIDRespond struct {
	Result
}

// GetUsersByIDRespond only for api docs
type GetUsersByIDRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Users UsersObjDetail `json:"users"`
	} `json:"data"` // return data
}

// DeleteUsersByIDRespond only for api docs
type DeleteUsersByIDRespond struct {
	Result
}

// DeleteUserssByIDsRequest request params
type DeleteUserssByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// DeleteUserssByIDsRespond only for api docs
type DeleteUserssByIDsRespond struct {
	Result
}

// GetUsersByConditionRequest request params
type GetUsersByConditionRequest struct {
	query.Conditions
}

// GetUsersByConditionRespond only for api docs
type GetUsersByConditionRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Users UsersObjDetail `json:"users"`
	} `json:"data"` // return data
}

// ListUserssByIDsRequest request params
type ListUserssByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// ListUserssByIDsRespond only for api docs
type ListUserssByIDsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Userss []UsersObjDetail `json:"userss"`
	} `json:"data"` // return data
}

// ListUserssRequest request params
type ListUserssRequest struct {
	query.Params
}

// ListUserssRespond only for api docs
type ListUserssRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Userss []UsersObjDetail `json:"userss"`
	} `json:"data"` // return data
}
