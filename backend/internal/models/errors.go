package models

import "errors"

var (
	StatusOK    = 0
	StatusError = -1
)
var (
	ErrInvalidInput   = errors.New("INVALID_INPUT")
	ErrInternalServer = errors.New("INTERNAL_SERVER_ERROR")
	ErrNoRows         = errors.New("NO_FLIGHT")
	ErrUnauthorized   = errors.New("UNAUTHORIZED")
	ErrForbidden      = errors.New("FORBIDDEN")
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
