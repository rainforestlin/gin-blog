package service

import (
	"github.com/julianlee107/blogWithGin/internal/model"
	"github.com/julianlee107/blogWithGin/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type GetTagRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,defalut=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name string `form:"name" binding:"max=100"`
	// 必须是0或1，如果不是默认为1
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,max=100,min=1"`
	CreatedBy string `form:"created_by" binding:"required,max=100,min=1"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"required,max=100,min=1"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,max=100,min=1"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int64, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

// 获取单个tag
func (svc *Service) GetSingleTag(param *GetTagRequest) (*model.Tag, error) {
	return svc.dao.GetTag(param.ID, param.State)
}

// 获取多个tag
func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.CreatedBy, param.State)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
