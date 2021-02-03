package dao

import "github.com/julianlee107/blogWithGin/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	return model.Auth{AppKey: appKey, AppSecret: appSecret}.Get(d.engine)
}
