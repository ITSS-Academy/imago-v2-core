package category

import (
	"context"
	"errors"
	"github.com/itss-academy/imago/core/common"
)

type Category struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Icon  string   `json:"icon"`
	Users []string `json:"users"`
}

type CategoryRepository interface {
	Create(ctx context.Context, category *Category) error
	GetById(ctx context.Context, id string) (*Category, error)
	GetByPage(ctx context.Context, opts *common.QueryOpts) ([]*Category, error)
	Get(ctx context.Context) ([]*Category, error)
	Update(ctx context.Context, category *Category) error
}

type CategoryUseCase interface {
	Create(ctx context.Context, category *Category) error
	GetById(ctx context.Context, id string) (*Category, error)
	GetByPage(ctx context.Context, opts *common.QueryOpts) ([]*Category, error)
	Get(ctx context.Context) ([]*Category, error)
	Update(ctx context.Context, category *Category) error
}

type CategoryInterop interface {
	Create(ctx context.Context, category *Category) error
	GetById(ctx context.Context, id string) (*Category, error)
	GetByPage(ctx context.Context, opts *common.QueryOpts) ([]*Category, error)
	Get(ctx context.Context) ([]*Category, error)
	Update(ctx context.Context, category *Category) error
}

var (
	ErrNameEmpty          = errors.New("name is empty")
	ErrIconEmpty          = errors.New("icon is empty")
	ErrIDEmpty            = errors.New("id is empty")
	ErrCategoryNotCreated = errors.New("category is not created")
	ErrCategoryNotFound   = errors.New("category is not found")
	ErrCategoryNotUpdated = errors.New("can not update category")
)
