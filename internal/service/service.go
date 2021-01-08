package service

import (
	"context"

	"github.com/julianlee107/blogWithGin/global"
	"github.com/julianlee107/blogWithGin/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	return Service{
		ctx: ctx,
		dao: dao.New(global.DBEngine),
	}
}
