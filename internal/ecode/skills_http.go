package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// skills business-level http error codes.
// the skillsNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	skillsNO       = 71
	skillsName     = "skills"
	skillsBaseCode = errcode.HCode(skillsNO)

	ErrCreateSkills         = errcode.NewError(skillsBaseCode+1, "failed to create "+skillsName)
	ErrDeleteByIDSkills     = errcode.NewError(skillsBaseCode+2, "failed to delete "+skillsName)
	ErrDeleteByIDsSkills    = errcode.NewError(skillsBaseCode+3, "failed to delete by batch ids "+skillsName)
	ErrUpdateByIDSkills     = errcode.NewError(skillsBaseCode+4, "failed to update "+skillsName)
	ErrGetByIDSkills        = errcode.NewError(skillsBaseCode+5, "failed to get "+skillsName+" details")
	ErrGetByConditionSkills = errcode.NewError(skillsBaseCode+6, "failed to get "+skillsName+" details by conditions")
	ErrListByIDsSkills      = errcode.NewError(skillsBaseCode+7, "failed to list by batch ids "+skillsName)
	ErrListByLastIDSkills   = errcode.NewError(skillsBaseCode+8, "failed to list by last id "+skillsName)
	ErrListSkills           = errcode.NewError(skillsBaseCode+9, "failed to list of "+skillsName)
	// error codes are globally unique, adding 1 to the previous error code
)
