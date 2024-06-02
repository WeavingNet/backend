package types

import (
	"time"

	"github.com/zhufuyi/sponge/pkg/ggorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateSkillsRequest request params
type CreateSkillsRequest struct {
	UserID           int    `json:"userId" binding:""`           // 用户ID
	SkillType        string `json:"skillType" binding:""`        // 技能类型
	SkillName        string `json:"skillName" binding:""`        // 技能名称
	ProficiencyLevel string `json:"proficiencyLevel" binding:""` // 熟练程度
}

// UpdateSkillsByIDRequest request params
type UpdateSkillsByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id

	UserID           int    `json:"userId" binding:""`           // 用户ID
	SkillType        string `json:"skillType" binding:""`        // 技能类型
	SkillName        string `json:"skillName" binding:""`        // 技能名称
	ProficiencyLevel string `json:"proficiencyLevel" binding:""` // 熟练程度
}

// SkillsObjDetail detail
type SkillsObjDetail struct {
	ID string `json:"id"` // convert to string id

	UserID           int       `json:"userId"`           // 用户ID
	SkillType        string    `json:"skillType"`        // 技能类型
	SkillName        string    `json:"skillName"`        // 技能名称
	ProficiencyLevel string    `json:"proficiencyLevel"` // 熟练程度
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

// CreateSkillsRespond only for api docs
type CreateSkillsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// UpdateSkillsByIDRespond only for api docs
type UpdateSkillsByIDRespond struct {
	Result
}

// GetSkillsByIDRespond only for api docs
type GetSkillsByIDRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Skills SkillsObjDetail `json:"skills"`
	} `json:"data"` // return data
}

// DeleteSkillsByIDRespond only for api docs
type DeleteSkillsByIDRespond struct {
	Result
}

// DeleteSkillssByIDsRequest request params
type DeleteSkillssByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// DeleteSkillssByIDsRespond only for api docs
type DeleteSkillssByIDsRespond struct {
	Result
}

// GetSkillsByConditionRequest request params
type GetSkillsByConditionRequest struct {
	query.Conditions
}

// GetSkillsByConditionRespond only for api docs
type GetSkillsByConditionRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Skills SkillsObjDetail `json:"skills"`
	} `json:"data"` // return data
}

// ListSkillssByIDsRequest request params
type ListSkillssByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// ListSkillssByIDsRespond only for api docs
type ListSkillssByIDsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Skillss []SkillsObjDetail `json:"skillss"`
	} `json:"data"` // return data
}

// ListSkillssRequest request params
type ListSkillssRequest struct {
	query.Params
}

// ListSkillssRespond only for api docs
type ListSkillssRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Skillss []SkillsObjDetail `json:"skillss"`
	} `json:"data"` // return data
}
