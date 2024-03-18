package repo

import (
	"context"
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

func (a AuthRepository) Get(ctx context.Context, opts *common.QueryOpts) ([]*auth.Auth, error) {
	authList := make([]*auth.Auth, 0)
	offset := opts.Page * opts.Size
	tx := a.db.Model(&auth.Auth{}).Offset(offset).Limit(opts.Size).Find(&authList)
	return authList, tx.Error
}

func (a AuthRepository) Update(ctx context.Context, auth *auth.Auth) error {
	tx := a.db.Save(auth)
	return tx.Error
}

func (a AuthRepository) Delete(ctx context.Context, id string) error {
	tx := a.db.Where("id = ?", id).Delete(&auth.Auth{})
	return tx.Error
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	err := db.AutoMigrate(&auth.Auth{})
	if err != nil {
		panic(err)
	}
	return &AuthRepository{db: db}
}
