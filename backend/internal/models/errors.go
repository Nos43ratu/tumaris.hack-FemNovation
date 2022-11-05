package models

import "errors"

var (
	StatusOK    = 0
	StatusError = -1
)
var (
	ErrInvalidInput       = errors.New("INVALID_INPUT")
	ErrInternalServer     = errors.New("INTERNAL_SERVER_ERROR")
	ErrNoRows             = errors.New("NO_FLIGHT")
	ErrUnauthorized       = errors.New("UNAUTHORIZED")
	ErrForbidden          = errors.New("FORBIDDEN")
	ErrUserDoesNotExist   = errors.New("USER_DOES_NOT_EXIST")
	ErrUserAlreadyExists  = errors.New("USER_ALREADY_EXISTS")
	ErrDBConnection       = errors.New("DB_CONNECTION_ERROR")
	ErrOrgNotFound        = errors.New("ORGANIZATION_NOT_FOUND")
	ErrEmailFormat        = errors.New("WRONG_EMAIL_FORMAT")
	ErrEmailServiceError  = errors.New("EMAIL_SERVICE_ERROR")
	ErrUserGroupNotFound  = errors.New("USER_GROUP_NOT_FOUND")
	ErrInviteDoesNotExist = errors.New("INVITE_DOES_NOT_EXIST")
	ErrWrongVerifyCode    = errors.New("WRONG_VERIFY_CODE")
)

type Answer struct {
	Response *Resp `json:"response"`
	Error    *Err  `json:"error"`
}

type Resp struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type Err struct {
	Message string `json:"message"`
}
