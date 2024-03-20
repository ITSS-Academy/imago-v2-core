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
	profileList := make([]*profile.Profile, 0)
	tx := p.db.Find(&profileList)
	return profileList, tx.Error
}

func (p ProfileRepository) Create(ctx context.Context, profileData *profile.Profile) error {
	tx := p.db.Create(profileData)
	return tx.Error
}

func (p ProfileRepository) Update(ctx context.Context, profile *profile.Profile) error {
	tx := p.db.Save(profile)
	return tx.Error
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	err := db.AutoMigrate(&profile.Profile{})
	if err != nil {
		panic(err)
	}
	return &ProfileRepository{db: db}
}
