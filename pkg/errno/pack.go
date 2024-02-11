package errno

import (
	"bibi/biz/model/interaction"
	"bibi/biz/model/user"
	"bibi/biz/model/video"
	"errors"
)

func BuildUserBaseResp(err error) *user.BaseResp {
	if err == nil {
		return ErrToUserResp(Success)
	}
	e := ErrNo{}
	if errors.As(err, &e) {
		return ErrToUserResp(e)
	}
	_e := ServiceError.WithMessage(err.Error())
	return ErrToUserResp(_e)
}

func ErrToUserResp(err ErrNo) *user.BaseResp {
	return &user.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}

func BuildVideoBaseResp(err error) *video.BaseResp {
	if err == nil {
		return ErrToVideoResp(Success)
	}
	e := ErrNo{}
	if errors.As(err, &e) {
		return ErrToVideoResp(e)
	}
	_e := ServiceError.WithMessage(err.Error())
	return ErrToVideoResp(_e)
}

func ErrToVideoResp(err ErrNo) *video.BaseResp {
	return &video.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}

func IsAllowExt(fileExt string, allowExtMap map[string]bool) bool {
	if _, ok := allowExtMap[fileExt]; !ok {
		return false
	}
	return true
}

func BuildInteractionBaseResp(err error) *interaction.BaseResp {
	if err == nil {
		return ErrToInteractionResp(Success)
	}
	e := ErrNo{}
	if errors.As(err, &e) {
		return ErrToInteractionResp(e)
	}
	_e := ServiceError.WithMessage(err.Error())
	return ErrToInteractionResp(_e)
}

func ErrToInteractionResp(err ErrNo) *interaction.BaseResp {
	return &interaction.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
