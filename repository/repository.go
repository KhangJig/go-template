package repository

import (
	"context"
	"demo-service/repository/demo"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	DemoRepo demo.IDemo
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		DemoRepo: demo.New(getClient),
	}
}
