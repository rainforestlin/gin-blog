package dao

import "github.com/julianlee107/blogWithGin/internal/model"

func (d *Dao) CountTag(name string, state uint8) (int64, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTag(id uint32, state uint8) (*model.Tag, error) {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
		State: state,
	}
	return tag.Get(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}

	return tag.List(d.engine, page, pageSize)
}

func (d *Dao) CreateTag(name, createdBy string, state uint8) error {
	tag := model.Tag{
		Name: name,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
		State: state,
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, ModifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{
			ModifiedBy: ModifiedBy,
			ID:         id,
		},
		Name:  name,
		State: state,
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": ModifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}
	return tag.Delete(d.engine)
}
