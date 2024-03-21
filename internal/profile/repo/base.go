package repo

import (
	"context"
	"github.com/itss-academy/imago/core/domain/profile"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func (p ProfileRepository) GetById(ctx context.Context, id string) (*profile.Profile, error) {
	profileData := &profile.Profile{}
	tx := p.db.Where("id = ?", id).First(profileData)
	return profileData, tx.Error
}

func (p ProfileRepository) GetAll(ctx context.Context) ([]*profile.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileRepository) Create(ctx context.Context, profileData *profile.Profile) error {
	tx := p.db.Create(profileData)
	return tx.Error
}

func (p ProfileRepository) Update(ctx context.Context, profile *profile.Profile) error {
	//TODO implement me
	panic("implement me")
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	err := db.AutoMigrate(&profile.Profile{})
	if err != nil {
		panic(err)
	}
	return &ProfileRepository{db: db}
}
