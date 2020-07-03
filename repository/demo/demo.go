package demo

import (
	"context"
	"demo-service/model"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Demo struct {
	getDB func(ctx context.Context) *gorm.DB
}

func New(getDB func(ctx context.Context) *gorm.DB) IDemo {
	return &Demo{getDB}
}

func (p Demo) GetTableName(ctx context.Context) string {
	return p.getDB(ctx).NewScope(model.Demo{}).GetModelStruct().TableName(p.getDB(ctx))
}

func (d *Demo) Test(ctx context.Context) string {
	return "test"
}

func (p Demo) GetByID(ctx context.Context, id int64) (*model.Demo, error) {
	var obj model.Demo
	err := p.getDB(ctx).First(&obj, id).Error

	return &obj, errors.Wrap(err, "GetByID fail")
}
