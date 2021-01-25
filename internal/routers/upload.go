package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/julianlee107/blogWithGin/global"
	"github.com/julianlee107/blogWithGin/internal/service"
	"github.com/julianlee107/blogWithGin/pkg/app"
	"github.com/julianlee107/blogWithGin/pkg/convert"
	"github.com/julianlee107/blogWithGin/pkg/errcode"
	"github.com/julianlee107/blogWithGin/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()

	if err != nil {
		errResp := errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errResp)
		return
	}
	if fileHeader == nil || fileType < 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())

	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)

	if err != nil {
		global.Logger.Error("svc.UploadFile err:", err)
		errResp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errResp)
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
