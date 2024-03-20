package repo

import (
	"context"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/category"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func (c CategoryRepository) Create(ctx context.Context, category *category.Category) error {
	tx := c.db.Create(category)
	return tx.Error
}

func (c CategoryRepository) GetById(ctx context.Context, id string) (*category.Category, error) {
	data := &category.Category{}
	tx := c.db.Where("id = ?", id).First(data)
	return data, tx.Error
}

func (c CategoryRepository) GetByPage(ctx context.Context, opts *common.QueryOpts) ([]*category.Category, error) {
	list := make([]*category.Category, 0)
	offset := opts.Page * opts.Size
	tx := c.db.Model(&category.Category{}).Offset(offset).Limit(opts.Size).Find(&list)
	return list, tx.Error
}

func (c CategoryRepository) Get(ctx context.Context) ([]*category.Category, error) {
	data := make([]*category.Category, 0)
	tx := c.db.Find(&data)
	return data, tx.Error
}

func (c CategoryRepository) Update(ctx context.Context, category *category.Category) error {
	tx := c.db.Save(category)
	return tx.Error
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	err := db.AutoMigrate(&category.Category{})
	if err != nil {
		panic(err)
	}
	return &CategoryRepository{db: db}
}
