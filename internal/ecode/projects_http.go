package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// projects business-level http error codes.
// the projectsNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	projectsNO       = 63
	projectsName     = "projects"
	projectsBaseCode = errcode.HCode(projectsNO)

	ErrCreateProjects         = errcode.NewError(projectsBaseCode+1, "failed to create "+projectsName)
	ErrDeleteByIDProjects     = errcode.NewError(projectsBaseCode+2, "failed to delete "+projectsName)
	ErrDeleteByIDsProjects    = errcode.NewError(projectsBaseCode+3, "failed to delete by batch ids "+projectsName)
	ErrUpdateByIDProjects     = errcode.NewError(projectsBaseCode+4, "failed to update "+projectsName)
	ErrGetByIDProjects        = errcode.NewError(projectsBaseCode+5, "failed to get "+projectsName+" details")
	ErrGetByConditionProjects = errcode.NewError(projectsBaseCode+6, "failed to get "+projectsName+" details by conditions")
	ErrListByIDsProjects      = errcode.NewError(projectsBaseCode+7, "failed to list by batch ids "+projectsName)
	ErrListByLastIDProjects   = errcode.NewError(projectsBaseCode+8, "failed to list by last id "+projectsName)
	ErrListProjects           = errcode.NewError(projectsBaseCode+9, "failed to list of "+projectsName)
	// error codes are globally unique, adding 1 to the previous error code
)
