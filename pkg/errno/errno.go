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
	ReadFileErrCode
	UploadFileErrCode
	LikeExistErrCode
	LikeNotExistErrCode
	VideoNotExistErrCode
	CommentIsNotExistErrCode
	FollowExistErrCode
	FollowNotExistErrCode
	FollowMyselfErrCode
)

const (
	SuccessMsg               = "Success"
	ServerErrMsg             = "Service is unable to start successfully"
	ParamErrMsg              = "Wrong Parameter has been given"
	UserAlreadyExistErrMsg   = "User existed"
	UserIsNotExistErrMsg     = "User is not exist"
	PasswordIsNotVerifiedMsg = "Username or password not verified"
	AuthErrMsg               = "It is not your account"
	ReadFileErrMsg           = "Error when read file"
	UploadFileErrMsg         = "Upload file error"
	LikeExistErrMsg          = "You have liked this video"
	LikeNotExistErrMsg       = "You don't like this video"
	LikeActionErrMsg         = "favorite add failed"
	FollowExistErrMsg        = "you have followed"
	FollowNotExistErrMsg     = "you haven't followed"
	FollowMyselfErrMsg       = "you can't follow yourself"

	MessageAddFailedErrMsg    = "message add failed"
	FriendListNoPermissionMsg = "You can't query his friend list"
	VideoNotExistErrMsg       = "video is not exist"
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

	ExistUserError         = NewErrNo(ExistUserErrCode, UserAlreadyExistErrMsg)
	NotExistUserError      = NewErrNo(NotExistUserErrCode, UserIsNotExistErrMsg)
	PwdError               = NewErrNo(AuthFailedErrCode, PasswordIsNotVerifiedMsg)
	AuthorizationError     = NewErrNo(AuthFailedErrCode, AuthErrMsg)
	UploadFileError        = NewErrNo(UploadFileErrCode, UploadFileErrMsg)
	ReadFileError          = NewErrNo(ReadFileErrCode, ReadFileErrMsg)
	LikeExistError         = NewErrNo(LikeExistErrCode, LikeExistErrMsg)
	LikeNotExistError      = NewErrNo(LikeNotExistErrCode, LikeNotExistErrMsg)
	VideoNotExistError     = NewErrNo(VideoNotExistErrCode, VideoNotExistErrMsg)
	CommentIsNotExistError = NewErrNo(CommentIsNotExistErrCode, CommentIsNotExistErrMsg)
	FollowExistError       = NewErrNo(FollowExistErrCode, FollowExistErrMsg)
	FollowNotExistError    = NewErrNo(FollowNotExistErrCode, FollowNotExistErrMsg)
	FollowMyselfError      = NewErrNo(FollowMyselfErrCode, FollowMyselfErrMsg)
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
