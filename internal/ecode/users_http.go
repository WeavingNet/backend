package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// users business-level http error codes.
// the usersNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	usersNO       = 27
	usersName     = "users"
	usersBaseCode = errcode.HCode(usersNO)

	ErrCreateUsers         = errcode.NewError(usersBaseCode+1, "failed to create "+usersName)
	ErrDeleteByIDUsers     = errcode.NewError(usersBaseCode+2, "failed to delete "+usersName)
	ErrDeleteByIDsUsers    = errcode.NewError(usersBaseCode+3, "failed to delete by batch ids "+usersName)
	ErrUpdateByIDUsers     = errcode.NewError(usersBaseCode+4, "failed to update "+usersName)
	ErrGetByIDUsers        = errcode.NewError(usersBaseCode+5, "failed to get "+usersName+" details")
	ErrGetByConditionUsers = errcode.NewError(usersBaseCode+6, "failed to get "+usersName+" details by conditions")
	ErrListByIDsUsers      = errcode.NewError(usersBaseCode+7, "failed to list by batch ids "+usersName)
	ErrListByLastIDUsers   = errcode.NewError(usersBaseCode+8, "failed to list by last id "+usersName)
	ErrListUsers           = errcode.NewError(usersBaseCode+9, "failed to list of "+usersName)
	// error codes are globally unique, adding 1 to the previous error code
)
