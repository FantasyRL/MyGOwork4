package errno

import (
	"bibi/biz/model/user"
	"errors"
)

func BuildBaseResp(err error) *user.BaseResp {
	if err == nil {
		return ErrToResp(Success)
	}
	e := ErrNo{}
	if errors.As(err, &e) {
		return ErrToResp(e)
	}
	_e := ServiceError.WithMessage(err.Error())
	return ErrToResp(_e)
}

func ErrToResp(err ErrNo) *user.BaseResp {
	return &user.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
