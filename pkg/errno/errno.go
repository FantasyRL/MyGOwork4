package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode    = 10000
	ServiceErrCode = iota + 10000 //未知服务错误
	ParamErrCode                  //参数错误

	ExistUserErrCode
	NotExistUserErrCode
	AuthFailedErrCode //Authorization错误
)

const (
	SuccessMsg               = "Success"
	ServerErrMsg             = "Service is unable to start successfully"
	ParamErrMsg              = "Wrong Parameter has been given"
	UserAlreadyExistErrMsg   = "User existed"
	UserIsNotExistErrMsg     = "user is not exist"
	PasswordIsNotVerifiedMsg = "username or password not verified"
	FavoriteActionErrMsg     = "favorite add failed"

	MessageAddFailedErrMsg    = "message add failed"
	FriendListNoPermissionMsg = "You can't query his friend list"
	VideoIsNotExistErrMsg     = "video is not exist"
	CommentIsNotExistErrMsg   = "comment is not exist"
)

type ErrNo struct {
	ErrorCode int64
	ErrorMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("%s", e.ErrorMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo { //出现不被定义的错误时
	e.ErrorMsg = msg
	return e
}

var (
	Success      = NewErrNo(SuccessCode, SuccessMsg)
	ServiceError = NewErrNo(ServiceErrCode, ServerErrMsg)
	ParamError   = NewErrNo(ParamErrCode, ParamErrMsg)

	ExistUserError    = NewErrNo(ExistUserErrCode, UserAlreadyExistErrMsg)
	NotExistUserError = NewErrNo(NotExistUserErrCode, UserIsNotExistErrMsg)
	PwdError          = NewErrNo(AuthFailedErrCode, PasswordIsNotVerifiedMsg)
)

// ConvertErr convert error to ErrNo
// in Default user ServiceErrCode
func ConvertErr(err error) ErrNo {
	errno := ErrNo{}
	if errors.As(err, &errno) {
		return errno
	}

	s := ServiceError
	s.ErrorMsg = err.Error()
	return s
}
