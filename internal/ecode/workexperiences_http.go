package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// workexperiences business-level http error codes.
// the workexperiencesNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	workexperiencesNO       = 42
	workexperiencesName     = "workexperiences"
	workexperiencesBaseCode = errcode.HCode(workexperiencesNO)

	ErrCreateWorkexperiences         = errcode.NewError(workexperiencesBaseCode+1, "failed to create "+workexperiencesName)
	ErrDeleteByIDWorkexperiences     = errcode.NewError(workexperiencesBaseCode+2, "failed to delete "+workexperiencesName)
	ErrDeleteByIDsWorkexperiences    = errcode.NewError(workexperiencesBaseCode+3, "failed to delete by batch ids "+workexperiencesName)
	ErrUpdateByIDWorkexperiences     = errcode.NewError(workexperiencesBaseCode+4, "failed to update "+workexperiencesName)
	ErrGetByIDWorkexperiences        = errcode.NewError(workexperiencesBaseCode+5, "failed to get "+workexperiencesName+" details")
	ErrGetByConditionWorkexperiences = errcode.NewError(workexperiencesBaseCode+6, "failed to get "+workexperiencesName+" details by conditions")
	ErrListByIDsWorkexperiences      = errcode.NewError(workexperiencesBaseCode+7, "failed to list by batch ids "+workexperiencesName)
	ErrListByLastIDWorkexperiences   = errcode.NewError(workexperiencesBaseCode+8, "failed to list by last id "+workexperiencesName)
	ErrListWorkexperiences           = errcode.NewError(workexperiencesBaseCode+9, "failed to list of "+workexperiencesName)
	// error codes are globally unique, adding 1 to the previous error code
)
