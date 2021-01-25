package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/julianlee107/blogWithGin/global"
	"github.com/julianlee107/blogWithGin/internal/service"
	"github.com/julianlee107/blogWithGin/pkg/app"
	"github.com/julianlee107/blogWithGin/pkg/convert"
	"github.com/julianlee107/blogWithGin/pkg/errcode"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取单个标签
// @Produce json
// @Param id path int true "标签id"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags/{id} [get]
func (t Tag) Get(c *gin.Context) {
	param := service.GetTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}

	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Error("app.BindAndValid err: ", errs)
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errResp)
		return
	}
	svc := service.New(c.Request.Context())

	tag, err := svc.GetSingleTag(&param)

	if err != nil {
		global.Logger.Error("svc.GetSingleTag err: ", errs)
		response.ToErrorResponse(errcode.ErrorGetTagFail)
		return
	}
	response.ToResponse(tag)
}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Emums(0,1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {array} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Error("app.BindAndValid err: ", errs)
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errResp)
		return
	}

	svc := service.New(c.Request.Context())

	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})

	if err != nil {
		global.Logger.Error("app.CountTag err: ", errs)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Error("svc.GetTagList err: ", errs)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponseList(tags, totalRows)
}

// @Summary 新增标签
// @Pruduce json
// @Param tag body service.CreateTagRequest true "标签"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{
		CreatedBy: c.PostForm("created_by"),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Error("app.BindAndValid err: ", errs)
		global.Logger.Debug("param:", param, " created_by:", c.PostForm("created_by"))
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errResp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)

	if err != nil {
		global.Logger.Error("svc.CreateTag err:", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{
		"msg": "创建成功",
	})
}

// @Summary 更改标签
// @Pruduce json
// @Param id path int true "标签ID"
// @Param tag body service.UpdateTagRequest true "标签"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Error("app.BindAndValid errs:", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Error("svc.UpdateTag err: ", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 删除标签
// @Produce json
// @Param id path int true "标签ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {}
