package pack

import (
	"bibi/pkg/errno"
	"fmt"
	"io"
	"mime/multipart"
	"time"
)

func FileToByte(file *multipart.FileHeader) ([]byte, error) {
	fileContent, err := file.Open()
	if err != nil {
		return nil, errno.ParamError
	}
	return io.ReadAll(fileContent)
}

func GenerateName(id int64) (string, string) {
	currentTime := time.Now()
	// 获取年月日和小时分钟
	year, month, day := currentTime.Date()
	hour, minute := currentTime.Hour(), currentTime.Minute()
	second := currentTime.Second()
	return fmt.Sprintf("%v_%d%02d%02d_%02d%02d%02d_video.mp4", id, year, month, day, hour, minute, second),
		fmt.Sprintf("%v_%d%02d%02d_%02d%02d%02d_cover.jpg", id, year, month, day, hour, minute, second)
}
