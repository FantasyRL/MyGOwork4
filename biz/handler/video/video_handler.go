// Code generated by hertz generator.

package video

import (
	"bibi/biz/video/dal/db"
	"bibi/biz/video/pack"
	"bibi/biz/video/service"
	"bibi/pkg/conf"
	"bibi/pkg/errno"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"path/filepath"

	video "bibi/biz/model/video"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// PutVideo .
// @Summary PutVideo
// @Description put video
// @Accept json/form
// @Produce json
// @Param video_file formData file true "视频文件"
// @Param title query string true "标题"
// @Param cover formData file true "视频封面"
// @Param token query string true "token"
// @router /bibi/video/upload [POST]
func PutVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.PutVideoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	videoFile, err := c.FormFile("video_file")
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	cover, err := c.FormFile("cover")
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(video.PutVideoResp)
	//文件类型判断
	videoExt := filepath.Ext(videoFile.Filename)
	allowExtVideo := map[string]bool{
		".mp4": true,
	}
	if !errno.IsAllowExt(videoExt, allowExtVideo) {
		resp.Base = errno.BuildVideoBaseResp(errno.ParamError)
		c.JSON(consts.StatusOK, resp)
		return
	}
	coverExt := filepath.Ext(cover.Filename)
	allowExtCover := map[string]bool{
		".jpg":  true,
		".png":  true,
		".jpeg": true,
	}
	if !errno.IsAllowExt(coverExt, allowExtCover) {
		resp.Base = errno.BuildVideoBaseResp(errno.ParamError)
		c.JSON(consts.StatusOK, resp)
		return
	}
	//jwt(mw)
	v, _ := c.Get("current_user_id")
	id := v.(int64)
	//名字生成
	videoName, coverName := pack.GenerateName(id)
	//开启并发
	var eg errgroup.Group
	eg.Go(func() error {
		coverByte, err := pack.FileToByte(cover)
		if err != nil {
			return errno.ReadFileError
		}
		err = service.NewVideoService(ctx).UploadCover(coverByte, coverName)
		if err != nil {
			return errno.UploadFileError
		}
		return nil
	})
	eg.Go(func() error {
		videoByte, err := pack.FileToByte(videoFile)
		if err != nil {
			return errno.ReadFileError
		}
		err = service.NewVideoService(ctx).UploadVideo(videoByte, videoName)
		if err != nil {
			return errno.UploadFileError
		}
		return nil
	})
	VideoReq := new(db.Video)
	eg.Go(func() error {
		videoUrl := fmt.Sprintf("%s/%s/video/%s", conf.OSSConf.EndPoint, conf.OSSConf.MainDirectory, videoName)
		coverUrl := fmt.Sprintf("%s/%s/video/%s", conf.OSSConf.EndPoint, conf.OSSConf.MainDirectory, coverName)
		VideoReq = &db.Video{
			Uid:      id,
			Title:    req.Title,
			PlayUrl:  videoUrl,
			CoverUrl: coverUrl,
		}
		_, err = service.NewVideoService(ctx).PutVideo(VideoReq)
		if err != nil {
			return err
		}
		return nil
	})
	if err := eg.Wait(); err != nil {
		resp.Base = errno.BuildVideoBaseResp(err)
		c.JSON(consts.StatusOK, resp)
		return
	}
	resp.Base = errno.BuildVideoBaseResp(nil)
	c.JSON(consts.StatusOK, resp)
}

// ListVideo .
// @router /bibi/video/myvideo [POST]
func ListVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.ListUserVideoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(video.ListUserVideoResp)

	c.JSON(consts.StatusOK, resp)
}

// SearchVideo .
// @router /bibi/video/search [POST]
func SearchVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.SearchVideoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(video.SearchVideoResp)

	c.JSON(consts.StatusOK, resp)
}

// HotVideo .
// @router /bibi/video/hot [GET]
func HotVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.HotVideoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(video.HotVideoReq)

	c.JSON(consts.StatusOK, resp)
}
