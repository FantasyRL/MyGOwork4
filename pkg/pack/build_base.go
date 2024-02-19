package pack

import (
	"bibi/biz/model/follow"
	"bibi/biz/model/interaction"
	"bibi/biz/model/user"
	"bibi/biz/model/video"
	"bibi/pkg/errno"
	"errors"
)

func BuildUserBaseResp(err error) *user.BaseResp {
	if err == nil {
		return ErrToUserResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return ErrToUserResp(e)
	}
	_e := errno.ServiceError.WithMessage(err.Error())
	return ErrToUserResp(_e)
}

func ErrToUserResp(err errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}

func BuildVideoBaseResp(err error) *video.BaseResp {
	if err == nil {
		return ErrToVideoResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return ErrToVideoResp(e)
	}
	_e := errno.ServiceError.WithMessage(err.Error())
	return ErrToVideoResp(_e)
}

func ErrToVideoResp(err errno.ErrNo) *video.BaseResp {
	return &video.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}

func BuildInteractionBaseResp(err error) *interaction.BaseResp {
	if err == nil {
		return ErrToInteractionResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return ErrToInteractionResp(e)
	}
	_e := errno.ServiceError.WithMessage(err.Error())
	return ErrToInteractionResp(_e)
}

func ErrToInteractionResp(err errno.ErrNo) *interaction.BaseResp {
	return &interaction.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}

func BuildFollowBaseResp(err error) *follow.BaseResp {
	if err == nil {
		return ErrToFollowResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return ErrToFollowResp(e)
	}
	_e := errno.ServiceError.WithMessage(err.Error())
	return ErrToFollowResp(_e)
}

func ErrToFollowResp(err errno.ErrNo) *follow.BaseResp {
	return &follow.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
