// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf2-demo/internal/model"
	"gf2-demo/internal/model/entity"
)

type (
	IDemo interface {
		Create(ctx context.Context, in model.DemoCreateInput) (*model.DemoCreateOutput, error)
		Update(ctx context.Context, in model.DemoUpdateInput) error
		GetInfo(ctx context.Context, fileda string) (*entity.Demo, error)
		Delete(ctx context.Context, id uint) error
		IDNotFound(ctx context.Context, id uint) (bool, error)
		FieldaNotFound(ctx context.Context, fielda string) (bool, error)
	}
)

var (
	localDemo IDemo
)

func Demo() IDemo {
	if localDemo == nil {
		panic("implement not found for interface IDemo, forgot register?")
	}
	return localDemo
}

func RegisterDemo(i IDemo) {
	localDemo = i
}
