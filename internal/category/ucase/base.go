package ucase

import (
	"context"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/category"
)

type CategoryUseCase struct {
	repo category.CategoryRepository
}

func (c CategoryUseCase) Create(ctx context.Context, categoryData *category.Category) error {
	err := c.Validate(categoryData)
	if err != nil {
		return err
	}
	err = c.repo.Create(ctx, categoryData)
	if err != nil {
		return category.ErrCategoryNotCreated
	}
	return nil
}

func (c CategoryUseCase) GetById(ctx context.Context, id string) (*category.Category, error) {
	data, err := c.repo.GetById(ctx, id)
	if err != nil {
		return nil, category.ErrCategoryNotFound
	}
	return data, nil
}

func (c CategoryUseCase) GetByPage(ctx context.Context, opts *common.QueryOpts) ([]*category.Category, error) {
	data, err := c.repo.GetByPage(ctx, opts)
	if err != nil {
		return nil, category.ErrCategoryNotFound
	}
	return data, nil
}

func (c CategoryUseCase) Get(ctx context.Context) ([]*category.Category, error) {
	data, err := c.repo.Get(ctx)
	if err != nil {
		return nil, category.ErrCategoryNotFound
	}
	return data, nil
}

func (c CategoryUseCase) Update(ctx context.Context, categoryData *category.Category) error {
	err := c.Validate(categoryData)
	if err != nil {
		return err
	}
	err = c.repo.Update(ctx, categoryData)
	if err != nil {
		return category.ErrCategoryNotUpdated
	}
	return nil
}

func (c CategoryUseCase) Validate(data *category.Category) error {
	if data.Name == "" {
		return category.ErrNameEmpty
	}
	if data.Icon == "" {
		return category.ErrIconEmpty
	}
	if data.ID == "" {
		return category.ErrIDEmpty
	}
	return nil
}

func NewCategoryUseCase(repo category.CategoryRepository) *CategoryUseCase {

	return &CategoryUseCase{repo: repo}
}
