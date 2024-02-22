// Code generated by hertz generator.
//MVC--View

package user

import (
	"bibi/biz/dal/db"
	"bibi/biz/model/user"
	"bibi/biz/service/user_service"
	"bibi/pkg/conf"
	"bibi/pkg/errno"
	"bibi/pkg/pack"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"golang.org/x/sync/errgroup"
	"path/filepath"
)

// Register .
// @Summary Register
// @Description userRegister
// @Accept json/form
// @Produce json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @router /bibi/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.RegisterResp)

	userResp, err := user_service.NewUserService(ctx).Register(&req)
	resp.Base = pack.BuildUserBaseResp(err)
	if err != nil {
		c.JSON(consts.StatusOK, resp)
		return
	}
	resp.UserID = userResp.ID
	c.JSON(consts.StatusOK, resp)
}

// Login .
// @Summary Login
// @Description userLogin
// @Accept json/form
// @Produce json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @router /bibi/user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	resp := new(user.LoginResp)

	resp.Base = pack.BuildUserBaseResp(nil)
	//hertz jwt(mw)
	v1, _ := c.Get("user")
	resp.User = user_service.BuildUserResp(v1)
	//hertz jwt(mw)
	v2, _ := c.Get("token")
	resp.Token = v2.(string)

	c.JSON(consts.StatusOK, resp)
}

// Info .
// @Summary Information
// @Description show user's info
// @Accept json/form
// @Produce json
// @Param Authorization header string true "token"
// @router /bibi/user [GET]
func Info(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.InfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.InfoResp)

	v, ok := c.Get("current_user_id")
	if !ok {
		err = errno.ParamError
	}
	id, _ := v.(int64)
	UserResp, err := user_service.NewUserService(ctx).Info(id)
	//hertz jwt(mw)

	resp.Base = pack.BuildUserBaseResp(err)
	if err != nil {
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.User = user_service.BuildUserResp(UserResp)
	c.JSON(consts.StatusOK, resp)
}

// Avatar .
// @Summary PUTAvatar
// @Description revise user's avatar
// @Accept json/form
// @Produce json
// @Param avatar_file formData file true "头像"
// @Param Authorization header string true "token"
// @router /bibi/user/avatar/upload [PUT]
func Avatar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.AvatarReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	//绑定文件
	file, err := c.FormFile("avatar_file")
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.AvatarResp)

	//判断文件格式
	fileExt := filepath.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".jpeg": true,
	}
	if !pack.IsAllowExt(fileExt, allowExtMap) {
		resp.Base = pack.BuildUserBaseResp(errno.ParamError)
		c.JSON(consts.StatusOK, resp)
		return
	}
	//hertz jwt(mw)
	v, _ := c.Get("current_user_id")
	id := v.(int64)
	//开启并发
	var eg errgroup.Group
	//上传至OSS
	eg.Go(func() error { //同时Add(1)
		req.AvatarFile, err = pack.FileToByte(file)
		if err != nil {
			return errno.ReadFileError
		}
		err = user_service.NewAvatarService(ctx).UploadAvatar(&req, id)
		if err != nil {
			return errno.UploadFileError
		}
		return nil
	})
	//上传url至数据库
	UserResp := new(db.User)
	eg.Go(func() error { //Add(1)
		avatarUrl := fmt.Sprintf("%s/%s/%d", conf.OSS.EndPoint, conf.OSS.MainDirectory, id)
		UserResp, err = user_service.NewAvatarService(ctx).PutAvatar(id, avatarUrl)
		if err != nil {
			return err
		}
		return nil
	})
	//Wait实现了错误处理与sync，仅返回第一个发生的错误
	if err := eg.Wait(); err != nil {
		resp.Base = pack.BuildUserBaseResp(err)
		c.JSON(consts.StatusOK, resp)
		return
	}
	resp.Base = pack.BuildUserBaseResp(nil)
	resp.User = user_service.BuildUserResp(UserResp)
	c.JSON(consts.StatusOK, resp)
}
