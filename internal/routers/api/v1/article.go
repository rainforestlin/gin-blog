package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/julianlee107/blogWithGin/global"
	"github.com/julianlee107/blogWithGin/internal/service"
	"github.com/julianlee107/blogWithGin/pkg/app"
	"github.com/julianlee107/blogWithGin/pkg/convert"
	"github.com/julianlee107/blogWithGin/pkg/errcode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary 获取单篇文章
// @Produce json
// @Param id path int true "标签id"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/article/{id} [get]
func (a Article) Get(c *gin.Context) {
	param := service.ArticleRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Error("app.BindAndValid errs:", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())

	article, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.Error("svc.GetArticle err:", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}
	response.ToResponse(article)
}

// @Summary 获取多篇文章
// @Produce json
// @Param title query string false "文章标题" maxlength(100)
// @Param state query int false "状态" Emums(0,1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {array} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Error("app.BindAndValid errs:", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	aricles, totalRows, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		global.Logger.Error("svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticlesFail)
		return
	}
	response.ToResponseList(aricles, totalRows)
}

// @Summary 新增文章
// @Pruduce json
// @Param title body string true "文章标题" minlength(1) maxlength(100)
// @Param state body int false "状态" Emums(0,1) default(1)
// @Param desc body string false "文章简介"
// @Param content body string false "文章内容"
// @Param created_by body string true "创建者" minlength(1) maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Error("app.BindAndValid errs:", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Error("svc.CreateArticle err:", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	response.ToResponse(gin.H{"msg": "create article successfully"})
}

// @Summary 修改文章
// @Pruduce json
// @Param title body string true "文章标题" minlength(1) maxlength(100)
// @Param state body int false "状态" Emums(0,1) default(1)
// @Param desc body string false "文章简介"
// @Param content body string false "文章内容"
// @Param modified_by body string true "创建者" minlength(1) maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Error("app.BindAndValid errs:", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Error("svc.UpdateArticle err:", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 删除单篇文章
// @Produce json
// @Param id path int true "标签id"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/article/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Error("app.BindAndValid errs:", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.Error("svc.DeleteArticle err:", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}
