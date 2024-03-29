package repo

import (
	"context"
	"math"

	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/auth"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func (a AuthRepository) Create(ctx context.Context, auth *auth.Auth) error {
	tx := a.db.Create(auth)
	return tx.Error
}

func (a AuthRepository) GetById(ctx context.Context, id string) (*auth.Auth, error) {
	authData := &auth.Auth{}
	tx := a.db.Where("id = ?", id).First(authData)
	return authData, tx.Error
}

func (a AuthRepository) Get(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*auth.Auth], error) {
	authData := make([]*auth.Auth, 0)
	limit := opts.Size
	offset := opts.Size * (opts.Page - 1)
	tx := a.db.Limit(limit).Offset(offset).Find(&authData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	count := int64(0)
	tx = a.db.Model(&auth.Auth{}).Count(&count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	pageNum := int(math.Ceil(float64(count) / float64(limit)))
	return &common.ListResult[*auth.Auth]{Data: authData, EndPage: int(pageNum)}, tx.Error
}

func (a AuthRepository) Update(ctx context.Context, auth *auth.Auth) error {
	tx := a.db.Save(auth)
	return tx.Error
}

func (a AuthRepository) Delete(ctx context.Context, id string) error {
	tx := a.db.WithContext(ctx).Where("id = ?", id).Delete(&auth.Auth{})
	return tx.Error
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	err := db.AutoMigrate(&auth.Auth{})
	if err != nil {
		panic(err)
	}
	return &AuthRepository{db: db}
}
