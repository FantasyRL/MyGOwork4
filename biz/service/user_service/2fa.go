package user_service

import (
	"bibi/biz/dal/db"
	"bibi/biz/model/user"
	"bibi/pkg/errno"
	"bibi/pkg/utils/otp2fa"
	"bibi/pkg/utils/sender"
	"bytes"
	"fmt"
	"image/png"
)

func (s *UserService) Switch2faType(req *user.Switch2FAReq, uid int64) error {
	userResp, err := db.QueryUserByID(&db.User{ID: uid})
	fmt.Println(*userResp)
	if err != nil {
		return err
	}
	//if req.ActionType == userResp.Type2fa {
	//	switch req.ActionType {
	//	case 1:
	//		return errno.Enable2FAError
	//	case 0:
	//		return errno.Unable2FAError
	//	}
	//}
	switch req.ActionType {
	case 1:
		key, err := otp2fa.GenerateTotp(userResp.Email)
		if err != nil {
			return err
		}

		if db.Update2FA(key.Secret(), uid) != nil {
			return err
		}

		qrcode, _ := key.Image(200, 200)
		buf := new(bytes.Buffer)
		_ = png.Encode(buf, qrcode)

		err = sender.SendEmail(userResp, buf)
		if err != nil {
			return err
		}
	case 0:
		if req.Totp == nil {
			return errno.ParamError
		}
		fmt.Println(*req.Totp)
		if !otp2fa.CheckTotp(*req.Totp, userResp.Otp) {
			return errno.Verify2FAError
		}
	}
	return db.Update2FAType(req.ActionType, uid)

}
