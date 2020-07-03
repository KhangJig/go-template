package demo

import (
	"context"
	"demo-service/model"
)

type IDemo interface {
	Test(ctx context.Context) string
	GetByID(ctx context.Context, id int64) (*model.Demo, error)
}
