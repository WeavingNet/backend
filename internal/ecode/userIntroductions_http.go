package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// userIntroductions business-level http error codes.
// the userIntroductionsNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	userIntroductionsNO       = 86
	userIntroductionsName     = "userIntroductions"
	userIntroductionsBaseCode = errcode.HCode(userIntroductionsNO)

	ErrCreateUserIntroductions         = errcode.NewError(userIntroductionsBaseCode+1, "failed to create "+userIntroductionsName)
	ErrDeleteByIDUserIntroductions     = errcode.NewError(userIntroductionsBaseCode+2, "failed to delete "+userIntroductionsName)
	ErrDeleteByIDsUserIntroductions    = errcode.NewError(userIntroductionsBaseCode+3, "failed to delete by batch ids "+userIntroductionsName)
	ErrUpdateByIDUserIntroductions     = errcode.NewError(userIntroductionsBaseCode+4, "failed to update "+userIntroductionsName)
	ErrGetByIDUserIntroductions        = errcode.NewError(userIntroductionsBaseCode+5, "failed to get "+userIntroductionsName+" details")
	ErrGetByConditionUserIntroductions = errcode.NewError(userIntroductionsBaseCode+6, "failed to get "+userIntroductionsName+" details by conditions")
	ErrListByIDsUserIntroductions      = errcode.NewError(userIntroductionsBaseCode+7, "failed to list by batch ids "+userIntroductionsName)
	ErrListByLastIDUserIntroductions   = errcode.NewError(userIntroductionsBaseCode+8, "failed to list by last id "+userIntroductionsName)
	ErrListUserIntroductions           = errcode.NewError(userIntroductionsBaseCode+9, "failed to list of "+userIntroductionsName)
	// error codes are globally unique, adding 1 to the previous error code
)
