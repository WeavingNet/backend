package types

import (
	"time"

	"github.com/zhufuyi/sponge/pkg/ggorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateUserIntroductionsRequest request params
type CreateUserIntroductionsRequest struct {
	UserID  int    `json:"userId" binding:""`
	Title   string `json:"title" binding:""`   // 介绍标题
	Content string `json:"content" binding:""` // 介绍内容
}

// UpdateUserIntroductionsByIDRequest request params
type UpdateUserIntroductionsByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id

	UserID  int    `json:"userId" binding:""`
	Title   string `json:"title" binding:""`   // 介绍标题
	Content string `json:"content" binding:""` // 介绍内容
}

// UserIntroductionsObjDetail detail
type UserIntroductionsObjDetail struct {
	ID string `json:"id"` // convert to string id

	UserID    int       `json:"userId"`
	Title     string    `json:"title"`   // 介绍标题
	Content   string    `json:"content"` // 介绍内容
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CreateUserIntroductionsRespond only for api docs
type CreateUserIntroductionsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// UpdateUserIntroductionsByIDRespond only for api docs
type UpdateUserIntroductionsByIDRespond struct {
	Result
}

// GetUserIntroductionsByIDRespond only for api docs
type GetUserIntroductionsByIDRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		UserIntroductions UserIntroductionsObjDetail `json:"userIntroductions"`
	} `json:"data"` // return data
}

// DeleteUserIntroductionsByIDRespond only for api docs
type DeleteUserIntroductionsByIDRespond struct {
	Result
}

// DeleteUserIntroductionssByIDsRequest request params
type DeleteUserIntroductionssByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// DeleteUserIntroductionssByIDsRespond only for api docs
type DeleteUserIntroductionssByIDsRespond struct {
	Result
}

// GetUserIntroductionsByConditionRequest request params
type GetUserIntroductionsByConditionRequest struct {
	query.Conditions
}

// GetUserIntroductionsByConditionRespond only for api docs
type GetUserIntroductionsByConditionRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		UserIntroductions UserIntroductionsObjDetail `json:"userIntroductions"`
	} `json:"data"` // return data
}

// ListUserIntroductionssByIDsRequest request params
type ListUserIntroductionssByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// ListUserIntroductionssByIDsRespond only for api docs
type ListUserIntroductionssByIDsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		UserIntroductionss []UserIntroductionsObjDetail `json:"userIntroductionss"`
	} `json:"data"` // return data
}

// ListUserIntroductionssRequest request params
type ListUserIntroductionssRequest struct {
	query.Params
}

// ListUserIntroductionssRespond only for api docs
type ListUserIntroductionssRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		UserIntroductionss []UserIntroductionsObjDetail `json:"userIntroductionss"`
	} `json:"data"` // return data
}
