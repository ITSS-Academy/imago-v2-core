package interop

import (
	"context"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/category"
)

type CategoryInterop struct {
	useCase category.CategoryUseCase
}

func (c CategoryInterop) Create(ctx context.Context, categoryData *category.Category) error {
	categoryData.Name = categoryData.ID
	err := c.useCase.Create(ctx, categoryData)
	if err != nil {
		return err
	}
	return nil
}

func (c CategoryInterop) GetById(ctx context.Context, id string) (*category.Category, error) {
	data, err := c.useCase.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c CategoryInterop) GetByPage(ctx context.Context, opts *common.QueryOpts) ([]*category.Category, error) {
	data, err := c.useCase.GetByPage(ctx, opts)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c CategoryInterop) Get(ctx context.Context) ([]*category.Category, error) {
	data, err := c.useCase.Get(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c CategoryInterop) Update(ctx context.Context, categoryData *category.Category) error {
	err := c.useCase.Update(ctx, categoryData)
	if err != nil {
		return err
	}
	return nil
}

func NewCategoryInterop(useCase category.CategoryUseCase) *CategoryInterop {
	return &CategoryInterop{useCase: useCase}
}
