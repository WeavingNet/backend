package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// educations business-level http error codes.
// the educationsNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	educationsNO       = 85
	educationsName     = "educations"
	educationsBaseCode = errcode.HCode(educationsNO)

	ErrCreateEducations         = errcode.NewError(educationsBaseCode+1, "failed to create "+educationsName)
	ErrDeleteByIDEducations     = errcode.NewError(educationsBaseCode+2, "failed to delete "+educationsName)
	ErrDeleteByIDsEducations    = errcode.NewError(educationsBaseCode+3, "failed to delete by batch ids "+educationsName)
	ErrUpdateByIDEducations     = errcode.NewError(educationsBaseCode+4, "failed to update "+educationsName)
	ErrGetByIDEducations        = errcode.NewError(educationsBaseCode+5, "failed to get "+educationsName+" details")
	ErrGetByConditionEducations = errcode.NewError(educationsBaseCode+6, "failed to get "+educationsName+" details by conditions")
	ErrListByIDsEducations      = errcode.NewError(educationsBaseCode+7, "failed to list by batch ids "+educationsName)
	ErrListByLastIDEducations   = errcode.NewError(educationsBaseCode+8, "failed to list by last id "+educationsName)
	ErrListEducations           = errcode.NewError(educationsBaseCode+9, "failed to list of "+educationsName)
	// error codes are globally unique, adding 1 to the previous error code
)
